package smartbft

import (
	"bytes"
	"context"
	"encoding/hex"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hyperledger/fabric/orderer/consensus/smartbft/types"

	"github.com/pkg/errors"

	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/orderer/common/cluster"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/orderer"
	"github.com/hyperledger/fabric/protos/utils"
	"google.golang.org/grpc"
)

// BlocksStreamPuller fetches stream of blocks from active committee members
type BlocksStreamPuller struct {
	WriteBlock        func(block *common.Block)
	CommitteeSize     func() int
	lastCommitteeSize int
	OnBlockCommit     func(block *common.Block)
	RetryTimeout      time.Duration
	FetchTimeout      time.Duration

	headerPullers []*headerPuller

	Endpoints []cluster.EndpointCriteria
	LastBlock *common.Block
	TLSCert   []byte
	Channel   string
	Signer    crypto.LocalSigner

	Logger *flogging.FabricLogger

	BlockVerifier BlockVerifier
	Dialer        cluster.Dialer

	abortStream          func()
	lock                 sync.RWMutex
	connectionLock       sync.RWMutex
	blockDeliverEndpoint cluster.EndpointCriteria
	blockDeliverIndex    int
	conn                 *grpc.ClientConn
	stream               orderer.AtomicBroadcast_DeliverClient
	stopFlag             int32
}

//go:generate mockery --name=AtomicBroadcastServer --case=underscore --output=mocks/
type AtomicBroadcastServer interface {
	orderer.AtomicBroadcastServer
}

//go:generate mockery --name=ImpatientStream --case=underscore --output=mocks/
type ImpatientStream interface {
	orderer.AtomicBroadcast_DeliverClient
	Abort()
}

//go:generate mockery --name=BlockVerifier --case=underscore --output=mocks/
type BlockVerifier interface {
	cluster.BlockVerifier
}

//go:generate mockery --name=LedgerWriter --case=underscore --output=mocks/
type LedgerWriter interface {
	cluster.LedgerWriter
}

type StreamCreator func(conn *grpc.ClientConn) (orderer.AtomicBroadcast_DeliverClient, error)

func (p *BlocksStreamPuller) ContinuouslyPullBlocks() {
	p.lastCommitteeSize = 0
	p.Logger.Infof("Pulling blocks until re-introduced into the committee")

	for !p.isStopped() {
		block, err := p.tryFetchBlocks()
		if err != nil {
			p.Logger.Errorf("Failed to pull next block, reason %s", err)
			time.Sleep(p.RetryTimeout)
			p.disconnect()
			continue
		}

		p.Logger.Infof("Writing block (%d) to the ledger", block.Header.Number)
		p.WriteBlock(block)
		p.LastBlock = block

		p.OnBlockCommit(block)
	}

	p.Logger.Infof("Exiting loop")
}

// Stop halts blocks stream puller from fetching blocks
func (p *BlocksStreamPuller) Stop() {
	if p.isStopped() {
		return
	}
	p.Logger.Debugf("Stopping BlockStreamPuller")
	defer p.Logger.Infof("BlockStreamPuller stopped")
	atomic.StoreInt32(&p.stopFlag, 1)
	p.disconnect()
}

func (p *BlocksStreamPuller) Initialize(endpoints []cluster.EndpointCriteria, lastBlock *common.Block) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.Endpoints = endpoints
	p.LastBlock = lastBlock
	atomic.StoreInt32(&p.stopFlag, 0)
}

func (p *BlocksStreamPuller) allocateHeaderPullers() {
	p.lock.RLock()
	endpoints := p.Endpoints
	blockDeliverIndex := p.blockDeliverIndex
	p.lock.RUnlock()

	for i, endpoint := range endpoints {
		if i == blockDeliverIndex {
			continue
		}

		logger := p.Logger.With("endpoint", endpoint.Endpoint)

		hp := &headerPuller{
			blockVerifier:     p.BlockVerifier,
			committeeSize:     p.CommitteeSize,
			lastCommitteeSize: p.lastCommitteeSize,
			lastPulledSeq:     p.LastBlock.Header.Number,
			channel:           p.Channel,
			Signer:            p.Signer,
			tlsCert:           p.TLSCert,
			logger:            logger,
			retryInterval:     p.RetryTimeout,
			endpoint:          endpoint,
			dialer:            p.Dialer,
		}
		p.headerPullers = append(p.headerPullers, hp)
		hp.start()
	}
}

func (p *BlocksStreamPuller) detectCensorship(lastPulledSeq uint64) bool {
	n := len(p.headerPullers) + 1
	f := (n - 1) / 3

	var passedCount int

	p.Logger.Debugf("Our last pulled block sequence is %d", lastPulledSeq)
	for _, hp := range p.headerPullers {
		if hp.passedSeq(lastPulledSeq) {
			passedCount++
		}
	}

	return passedCount > f
}

// assignEndpoint assign endpoints to decide where to fetch blocks from
func (p *BlocksStreamPuller) assignEndpoint() {
	p.lock.RLock()
	defer p.lock.RUnlock()
	p.blockDeliverIndex = rand.Intn(len(p.Endpoints))
	p.blockDeliverEndpoint = p.Endpoints[p.blockDeliverIndex]
	p.Logger.Infof("Will pull from %s", p.blockDeliverEndpoint.Endpoint)
}

func (p *BlocksStreamPuller) tryFetchBlocks() (*common.Block, error) {
	if p.disconnected() {
		if err := p.obtainStream(); err != nil {
			p.Logger.Errorf("failed to obtain grpc stream to read blocks, because of %s", err)
			return nil, err
		}
		err := p.requestBlocks()
		if err != nil {
			p.Logger.Errorf("failed sending seek envelope, error %s, endpoint %s", err, p.blockDeliverEndpoint)
			return nil, err
		}

		p.allocateHeaderPullers()
	}

	noCensorship := make(chan struct{})

	go func() {
		for {
			timer := time.NewTimer(time.Second * 5)

			select {
			case <-timer.C:
				if p.detectCensorship(p.LastBlock.Header.Number) {
					p.Logger.Warnf("Detected censorship")
					p.abortStream()
					timer.Stop()
					return
				}
			case <-noCensorship:
				timer.Stop()
				return
			}
		}
	}()

	resp, err := p.stream.Recv()
	close(noCensorship)
	if err != nil {
		p.Logger.Errorf("Failed receiving next block from %s: %v", p.blockDeliverEndpoint, err)
		return nil, errors.Errorf("failed to receive next block, due to %s", err)
	}

	block, err := cluster.ExtractBlockFromResponse(resp)
	if err != nil {
		p.Logger.Errorf("Received a bad block from %s: %v", p.blockDeliverEndpoint, err)
		return nil, err
	}

	if p.LastBlock.Header.Number+1 != block.Header.Number {
		return nil, errors.Errorf("got unexpected sequence from %s - (%d) instead of (%d)", p.blockDeliverEndpoint.Endpoint, block.Header.Number, p.LastBlock.Header.Number+1)
	}
	if !bytes.Equal(block.Header.PreviousHash, p.LastBlock.Header.Hash()) {
		claimedPrevHash := hex.EncodeToString(block.Header.PreviousHash)
		actualPrevHash := hex.EncodeToString(p.LastBlock.Header.Hash())
		return nil,
			errors.Errorf("block [%d]'s hash (%s) mismatches %d's prev block hash (%s)",
				p.LastBlock.Header.Number, actualPrevHash, block.Header.Number, claimedPrevHash)
	}

	md, err := types.CommitteeMetadataFromBlock(block)
	if err != nil {
		return nil, errors.Wrapf(err, "failed extracting committee metadata from block %d", block.Header.Number)
	}

	committeeSize := p.lastCommitteeSize

	if committeeSize == 0 {
		committeeSize = p.CommitteeSize()
	}

	// sending nil for config envelope parameter, it will make use of the recent active configuration
	if err := cluster.VerifyBlockSignature(block, p.BlockVerifier, nil, committeeSize); err != nil {
		return nil, err
	}

	defer func() {
		if md != nil {
			p.lastCommitteeSize = int(md.CommitteeSize)
		}
	}()

	return block, nil
}

func (p *BlocksStreamPuller) newStream(conn *grpc.ClientConn) (orderer.AtomicBroadcast_DeliverClient, error) {
	abc := orderer.NewAtomicBroadcastClient(conn)
	ctx, cancel := context.WithCancel(context.Background())

	stream, err := abc.Deliver(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	p.connectionLock.Lock()
	defer p.connectionLock.Unlock()

	once := &sync.Once{}
	p.abortStream = func() {
		once.Do(func() {
			cancel()
		})
	}

	return stream, nil
}

func (p *BlocksStreamPuller) isStopped() bool {
	return atomic.LoadInt32(&p.stopFlag) == 1
}

// disconnect makes the BlockPuller close the connection and stream
// with the remote endpoint, and wipe the internal block buffer.
func (p *BlocksStreamPuller) disconnect() {
	p.connectionLock.Lock()
	defer p.connectionLock.Unlock()
	if p.stream != nil {
		p.abortStream()
	}

	if p.conn != nil {
		p.conn.Close()
	}
	p.conn = nil

	for _, hp := range p.headerPullers {
		hp.stop()
	}

	p.headerPullers = nil
}

func (p *BlocksStreamPuller) openStream() error {
	stream, err := p.newStream(p.conn)
	if err != nil {
		p.Logger.Errorf("failed to create impatient delivery stream to fetch blocks from committee, error %s", err)
		return err
	}
	p.stream = stream
	return nil
}

func (p *BlocksStreamPuller) obtainStream() error {
	reConnected := false
	for p.disconnected() {
		reConnected = true
		// not connected, need to select next point and reconnect to it
		// to continue fetching blocks
		p.assignEndpoint()
		// make sure to get connected to committee OSN
		if err := p.connectToNextEndpoint(); err != nil {
			time.Sleep(p.RetryTimeout)
		}
	}

	if reConnected {
		err := p.openStream()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *BlocksStreamPuller) requestBlocks() error {
	if p.stream == nil {
		return errors.Errorf("grps stream is not initialized, cannot pull blocks from committee")
	}
	env, err := utils.CreateSignedEnvelopeWithTLSBinding(
		common.HeaderType_DELIVER_SEEK_INFO,
		p.Channel,
		p.Signer,
		&orderer.SeekInfo{
			Start:         &orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: p.LastBlock.Header.Number + 1}}},
			Stop:          &orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: math.MaxUint64}}},
			Behavior:      orderer.SeekInfo_BLOCK_UNTIL_READY,
			ErrorResponse: orderer.SeekInfo_BEST_EFFORT,
		},
		int32(0),
		uint64(0),
		util.ComputeSHA256(p.TLSCert),
	)
	if err != nil {
		p.Logger.Errorf("failed to create seek envelope to start fetching blocks, endpoint %s", err, p.blockDeliverEndpoint)
		return err
	}
	// sending request to pull blocks
	p.Logger.Debugf("sending request to pull blocks from committee, endpoint is [%s]", p.blockDeliverEndpoint)
	return p.stream.Send(env)
}

func (p *BlocksStreamPuller) connectToNextEndpoint() error {
	var err error
	p.connectionLock.Lock()
	defer p.connectionLock.Unlock()
	p.conn, err = p.Dialer.Dial(p.blockDeliverEndpoint)
	if err != nil {
		p.Logger.Warningf("Failed connecting to %s: %v", p.blockDeliverEndpoint, err)
		return err
	}
	return nil
}

// disconnected check whenever puller got already connected
func (p *BlocksStreamPuller) disconnected() bool {
	p.connectionLock.RLock()
	defer p.connectionLock.RUnlock()

	return p.conn == nil
}

type headerPuller struct {
	// private
	running     sync.WaitGroup
	lock        sync.RWMutex
	closeChan   chan struct{}
	abortStream func()
	conn        *grpc.ClientConn
	// external
	committeeSize     func() int
	blockVerifier     BlockVerifier
	lastCommitteeSize int
	lastPulledSeq     uint64
	channel           string
	Signer            crypto.LocalSigner
	tlsCert           []byte
	logger            *flogging.FabricLogger
	retryInterval     time.Duration
	endpoint          cluster.EndpointCriteria
	dialer            cluster.Dialer
}

func (hp *headerPuller) stop() {
	defer hp.running.Wait()

	hp.lock.Lock()

	select {
	case <-hp.closeChan:
		hp.lock.Unlock()
		return
	default:
		close(hp.closeChan)
	}

	hp.lock.Unlock()

	hp.disconnect()

}

func (hp *headerPuller) start() {
	hp.closeChan = make(chan struct{})
	hp.running.Add(1)
	go func() {
		defer hp.running.Done()
		hp.run()
	}()
}

func (hp *headerPuller) closed() bool {
	hp.lock.Lock()
	defer hp.lock.Unlock()

	select {
	case <-hp.closeChan:
		return true
	default:
		return false
	}
}

func (hp *headerPuller) sleep() {
	select {
	case <-hp.closeChan:
		return
	case <-time.After(hp.retryInterval):
	}
}

func (hp *headerPuller) disconnect() {
	hp.lock.Lock()
	defer hp.lock.Unlock()

	if hp.abortStream != nil {
		hp.abortStream()
	}
	if hp.conn != nil {
		hp.conn.Close()
	}

	hp.conn = nil
	hp.abortStream = nil
}

func (hp *headerPuller) run() {
	for !hp.closed() {
		var err error
		hp.conn, err = hp.dialer.Dial(hp.endpoint)
		if err != nil {
			hp.logger.Warnf("Failed connecting to %s: %v", hp.endpoint.Endpoint, err)
			hp.sleep()
			continue
		}

		stream, err := hp.acquireStream()
		if err != nil {
			hp.disconnect()
			hp.logger.Warnf("Failed acquiring stream to %s: %v", hp.endpoint.Endpoint, err)
			hp.sleep()
			continue
		}

		if err := hp.requestBlocks(stream); err != nil {
			hp.disconnect()
			hp.logger.Warnf("Failed requesting block headers from %s: %v", hp.endpoint.Endpoint, err)
			continue
		}

		hp.pullHeaders(stream)

	}

}

func (hp *headerPuller) pullHeaders(stream orderer.AtomicBroadcast_DeliverClient) {
	for !hp.closed() {
		resp, err := stream.Recv()
		if err != nil {
			hp.logger.Warnf("Failed pulling block headers from %s: %v", hp.endpoint.Endpoint, err)
			hp.disconnect()
			hp.sleep()
			return
		}

		block := resp.GetBlock()
		if block == nil {
			hp.logger.Warnf("Deliver response we got from %s does not contain a block", hp.endpoint.Endpoint)
			hp.stop()
			return
		}

		if err := hp.verifyHeader(block); err != nil {
			hp.logger.Warnf("Failed verifying block header received from %s: %v", hp.endpoint.Endpoint, err)
			hp.sleep()
			hp.disconnect()
			return
		}

		hp.lock.Lock()
		hp.lastPulledSeq = block.Header.Number
		hp.lock.Unlock()
	}
}

func (hp *headerPuller) verifyHeader(block *common.Block) error {
	if block.Header == nil {
		return errors.Errorf("block header is nil")
	}
	if hp.lastPulledSeq+1 != block.Header.Number {
		return errors.Errorf("got unexpected sequence from %s - (%d) instead of (%d)", hp.endpoint.Endpoint, block.Header.Number, hp.lastPulledSeq+1)
	}

	md, err := types.CommitteeMetadataFromBlock(block)
	if err != nil {
		return errors.Wrapf(err, "failed extracting committee metadata from block %d", block.Header.Number)
	}

	committeeSize := hp.lastCommitteeSize

	if committeeSize == 0 {
		committeeSize = hp.committeeSize()
	}

	// sending nil for config envelope parameter, it will make use of the recent active configuration
	if err := cluster.VerifyBlockSignature(block, hp.blockVerifier, nil, committeeSize); err != nil {
		hp.logger.Warnf("Failed verifying block signature of block %d received from %s: %v", block.Header.Number, hp.endpoint.Endpoint, err)
		return err
	}

	defer func() {
		if md != nil {
			hp.lastCommitteeSize = int(md.CommitteeSize)
		}
	}()

	return nil
}

func (hp *headerPuller) passedSeq(lastPulledSeq uint64) bool {
	hp.lock.RLock()
	defer hp.lock.RUnlock()

	hp.logger.Debugf("Endpoint %s is at height %d", hp.endpoint.Endpoint, hp.lastPulledSeq)

	return hp.lastPulledSeq > lastPulledSeq
}

func (hp *headerPuller) acquireStream() (orderer.AtomicBroadcast_DeliverClient, error) {
	abc := orderer.NewAtomicBroadcastClient(hp.conn)
	ctx, cancel := context.WithCancel(context.Background())

	stream, err := abc.Deliver(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	hp.lock.Lock()
	defer hp.lock.Unlock()

	once := &sync.Once{}
	hp.abortStream = func() {
		once.Do(func() {
			cancel()
		})
	}

	return stream, nil
}

func (hp *headerPuller) requestBlocks(stream orderer.AtomicBroadcast_DeliverClient) error {
	env, err := utils.CreateSignedEnvelopeWithTLSBinding(
		common.HeaderType_DELIVER_SEEK_INFO,
		hp.channel,
		hp.Signer,
		&orderer.SeekInfo{
			Start:         &orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: hp.lastPulledSeq + 1}}},
			Stop:          &orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: math.MaxUint64}}},
			Behavior:      orderer.SeekInfo_BLOCK_UNTIL_READY,
			ErrorResponse: orderer.SeekInfo_BEST_EFFORT,
			ContentType:   orderer.SeekInfo_HEADER_WITH_SIG,
		},
		int32(0),
		uint64(0),
		util.ComputeSHA256(hp.tlsCert),
	)
	if err != nil {
		hp.logger.Errorf("failed to create seek envelope to start fetching blocks, endpoint %s", err, hp.endpoint)
		return err
	}
	hp.logger.Debugf("sending request to pull block headers from %s", hp.endpoint.Endpoint)
	return stream.Send(env)
}
