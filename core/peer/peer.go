/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package peer

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/orderer/smartbft"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/common/channelconfig"
	cc "github.com/hyperledger/fabric/common/config"
	"github.com/hyperledger/fabric/common/configtx"
	"github.com/hyperledger/fabric/common/deliver"
	"github.com/hyperledger/fabric/common/flogging"
	commonledger "github.com/hyperledger/fabric/common/ledger"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/common/semaphore"
	"github.com/hyperledger/fabric/core/committer"
	"github.com/hyperledger/fabric/core/committer/txvalidator"
	"github.com/hyperledger/fabric/core/committer/txvalidator/plugin"
	validatorv14 "github.com/hyperledger/fabric/core/committer/txvalidator/v14"
	validatorv20 "github.com/hyperledger/fabric/core/committer/txvalidator/v20"
	"github.com/hyperledger/fabric/core/committer/txvalidator/v20/plugindispatcher"
	vir "github.com/hyperledger/fabric/core/committer/txvalidator/v20/valinforetriever"
	"github.com/hyperledger/fabric/core/common/privdata"
	validation "github.com/hyperledger/fabric/core/handlers/validation/api/state"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/ledgermgmt"
	"github.com/hyperledger/fabric/core/transientstore"
	"github.com/hyperledger/fabric/gossip/api"
	gossipprivdata "github.com/hyperledger/fabric/gossip/privdata"
	gossipservice "github.com/hyperledger/fabric/gossip/service"
	"github.com/hyperledger/fabric/internal/pkg/comm"
	"github.com/hyperledger/fabric/internal/pkg/peer/orderers"
	"github.com/hyperledger/fabric/msp"
	mspmgmt "github.com/hyperledger/fabric/msp/mgmt"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/pkg/errors"
)

var peerLogger = flogging.MustGetLogger("peer")

type CollectionInfoShim struct {
	plugindispatcher.CollectionAndLifecycleResources
	ChannelID string
}

func (cis *CollectionInfoShim) CollectionValidationInfo(chaincodeName, collectionName string, validationState validation.State) ([]byte, error, error) {
	return cis.CollectionAndLifecycleResources.CollectionValidationInfo(cis.ChannelID, chaincodeName, collectionName, validationState)
}

type gossipSupport struct {
	channelconfig.Application
	configtx.Validator
	channelconfig.Channel
}

func ConfigBlockFromLedger(ledger ledger.PeerLedger) (*common.Block, error) {
	peerLogger.Debugf("Getting config block")

	// get last block.  Last block number is Height-1
	blockchainInfo, err := ledger.GetBlockchainInfo()
	if err != nil {
		return nil, err
	}
	lastBlock, err := ledger.GetBlockByNumber(blockchainInfo.Height - 1)
	if err != nil {
		return nil, err
	}

	// get most recent config block location from last block metadata
	configBlockIndex, err := protoutil.GetLastConfigIndexFromBlock(lastBlock)
	if err != nil {
		return nil, err
	}

	// get most recent config block
	configBlock, err := ledger.GetBlockByNumber(configBlockIndex)
	if err != nil {
		return nil, err
	}

	peerLogger.Debugf("Got config block[%d]", configBlockIndex)
	return configBlock, nil
}

// updates the trusted roots for the peer based on updates to channels
func (p *Peer) updateTrustedRoots(cm channelconfig.Resources) {
	if !p.ServerConfig.SecOpts.UseTLS {
		return
	}

	// this is triggered on per channel basis so first update the roots for the channel
	peerLogger.Debugf("Updating trusted root authorities for channel %s", cm.ConfigtxValidator().ChannelID())

	p.CredentialSupport.BuildTrustedRootsForChain(cm)

	// now iterate over all roots for all app and orderer channels
	var trustedRoots [][]byte
	for _, roots := range p.CredentialSupport.AppRootCAsByChain() {
		trustedRoots = append(trustedRoots, roots...)
	}
	trustedRoots = append(trustedRoots, p.ServerConfig.SecOpts.ClientRootCAs...)
	trustedRoots = append(trustedRoots, p.ServerConfig.SecOpts.ServerRootCAs...)

	// now update the client roots for the peerServer
	err := p.server.SetClientRootCAs(trustedRoots)
	if err != nil {
		msg := "Failed to update trusted roots from latest config block. " +
			"This peer may not be able to communicate with members of channel %s (%s)"
		peerLogger.Warningf(msg, cm.ConfigtxValidator().ChannelID(), err)
	}
}

//
//  Deliver service support structs for the peer
//

// DeliverChainManager provides access to a channel for performing deliver
type DeliverChainManager struct {
	Peer *Peer
}

func (d DeliverChainManager) GetChain(chainID string) deliver.Chain {
	if channel := d.Peer.Channel(chainID); channel != nil {
		return channel
	}
	return nil
}

// fileLedgerBlockStore implements the interface expected by
// common/ledger/blockledger/file to interact with a file ledger for deliver
type fileLedgerBlockStore struct {
	ledger.PeerLedger
}

func (flbs fileLedgerBlockStore) AddBlock(*common.Block) error {
	return nil
}

func (flbs fileLedgerBlockStore) RetrieveBlocks(startBlockNumber uint64) (commonledger.ResultsIterator, error) {
	return flbs.GetBlocksIterator(startBlockNumber)
}

func (flbs fileLedgerBlockStore) Shutdown() {}

// NewConfigSupport returns
func NewConfigSupport(peer *Peer) cc.Manager {
	return &configSupport{
		peer: peer,
	}
}

type configSupport struct {
	peer *Peer
}

// GetChannelConfig returns an instance of a object that represents
// current channel configuration tree of the specified channel. The
// ConfigProto method of the returned object can be used to get the
// proto representing the channel configuration.
func (c *configSupport) GetChannelConfig(cid string) cc.Config {
	channel := c.peer.Channel(cid)
	if channel == nil {
		peerLogger.Errorf("[channel %s] channel not associated with this peer", cid)
		return nil
	}
	return channel.Resources().ConfigtxValidator()
}

// A Peer holds references to subsystems and channels associated with a Fabric peer.
type Peer struct {
	ServerConfig             comm.ServerConfig
	CredentialSupport        *comm.CredentialSupport
	StoreProvider            transientstore.StoreProvider
	GossipService            *gossipservice.GossipService
	LedgerMgr                *ledgermgmt.LedgerMgr
	OrdererEndpointOverrides map[string]*orderers.Endpoint
	CryptoProvider           bccsp.BCCSP

	// validationWorkersSemaphore is used to limit the number of concurrent validation
	// go routines.
	validationWorkersSemaphore semaphore.Semaphore

	server             *comm.GRPCServer
	pluginMapper       plugin.Mapper
	channelInitializer func(cid string)

	// channels is a map of channelID to channel
	mutex    sync.RWMutex
	channels map[string]*Channel
}

func (p *Peer) openStore(cid string) (*transientstore.Store, error) {
	store, err := p.StoreProvider.OpenStore(cid)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (p *Peer) CreateChannel(
	cid string,
	cb *common.Block,
	deployedCCInfoProvider ledger.DeployedChaincodeInfoProvider,
	legacyLifecycleValidation plugindispatcher.LifecycleResources,
	newLifecycleValidation plugindispatcher.CollectionAndLifecycleResources,
) error {
	l, err := p.LedgerMgr.CreateLedger(cid, cb)
	if err != nil {
		return errors.WithMessage(err, "cannot create ledger from genesis block")
	}

	if err := p.createChannel(cid, l, deployedCCInfoProvider, legacyLifecycleValidation, newLifecycleValidation); err != nil {
		return err
	}

	p.initChannel(cid)
	return nil
}

// CreateChannelFromSnapshot creates a channel from the specified snapshot.
func (p *Peer) CreateChannelFromSnapshot(
	snapshotDir string,
	deployedCCInfoProvider ledger.DeployedChaincodeInfoProvider,
	legacyLifecycleValidation plugindispatcher.LifecycleResources,
	newLifecycleValidation plugindispatcher.CollectionAndLifecycleResources,
) error {
	channelCallback := func(l ledger.PeerLedger, cid string) {
		if err := p.createChannel(cid, l, deployedCCInfoProvider, legacyLifecycleValidation, newLifecycleValidation); err != nil {
			logger.Errorf("error creating channel for %s", cid)
			return
		}
		p.initChannel(cid)
	}

	err := p.LedgerMgr.CreateLedgerFromSnapshot(snapshotDir, channelCallback)
	if err != nil {
		return errors.WithMessagef(err, "cannot create ledger from snapshot %s", snapshotDir)
	}

	return nil
}

// RetrievePersistedChannelConfig retrieves the persisted channel config from statedb
func RetrievePersistedChannelConfig(ledger ledger.PeerLedger) (*common.Config, error) {
	qe, err := ledger.NewQueryExecutor()
	if err != nil {
		return nil, err
	}
	defer qe.Done()
	return retrieveChannelConfig(qe)
}

// SmartBFTId2Identities get identities from last known configuration.
func (p *Peer) SmartBFTId2Identities(cid string) map[uint64][]byte {
	c := p.Channel(cid)
	if c == nil {
		return nil
	}

	c.lock.RLock()
	defer c.lock.RUnlock()
	oc, ok := c.Resources().OrdererConfig()
	if !ok {
		return nil
	}

	m := &smartbft.ConfigMetadata{}
	proto.Unmarshal(oc.ConsensusMetadata(), m)

	res := make(map[uint64][]byte)
	for _, consenter := range m.Consenters {
		res[consenter.ConsenterId] = consenter.Identity
	}

	return res
}

type IdentityFethcer struct {
	Adaptee *Peer
}

func (i *IdentityFethcer) Id2Identities(cid string) map[uint64][]byte {
	return i.Adaptee.SmartBFTId2Identities(cid)
}

// createChannel creates a new channel object and insert it into the channels slice.
func (p *Peer) createChannel(
	cid string,
	l ledger.PeerLedger,
	deployedCCInfoProvider ledger.DeployedChaincodeInfoProvider,
	legacyLifecycleValidation plugindispatcher.LifecycleResources,
	newLifecycleValidation plugindispatcher.CollectionAndLifecycleResources,
) error {
	chanConf, err := RetrievePersistedChannelConfig(l)
	if err != nil {
		return err
	}

	bundle, err := channelconfig.NewBundle(cid, chanConf, p.CryptoProvider)
	if err != nil {
		return err
	}

	capabilitiesSupportedOrPanic(bundle)

	channelconfig.LogSanityChecks(bundle)

	gossipEventer := p.GossipService.NewConfigEventer()

	gossipCallbackWrapper := func(bundle *channelconfig.Bundle) {
		ac, ok := bundle.ApplicationConfig()
		if !ok {
			// TODO, handle a missing ApplicationConfig more gracefully
			ac = nil
		}
		gossipEventer.ProcessConfigUpdate(&gossipSupport{
			Validator:   bundle.ConfigtxValidator(),
			Application: ac,
			Channel:     bundle.ChannelConfig(),
		})
		p.GossipService.SuspectPeers(func(identity api.PeerIdentityType) bool {
			// TODO: this is a place-holder that would somehow make the MSP layer suspect
			// that a given certificate is revoked, or its intermediate CA is revoked.
			// In the meantime, before we have such an ability, we return true in order
			// to suspect ALL identities in order to validate all of them.
			return true
		})
	}

	trustedRootsCallbackWrapper := func(bundle *channelconfig.Bundle) {
		p.updateTrustedRoots(bundle)
	}

	mspCallback := func(bundle *channelconfig.Bundle) {
		// TODO remove once all references to mspmgmt are gone from peer code
		mspmgmt.XXXSetMSPManager(cid, bundle.MSPManager())
	}

	osLogger := flogging.MustGetLogger("peer.orderers")
	namedOSLogger := osLogger.With("channel", cid)
	ordererSource := orderers.NewConnectionSource(namedOSLogger, p.OrdererEndpointOverrides)

	ordererSourceCallback := func(bundle *channelconfig.Bundle) {
		globalAddresses := bundle.ChannelConfig().OrdererAddresses()
		orgAddresses := map[string]orderers.OrdererOrg{}
		if ordererConfig, ok := bundle.OrdererConfig(); ok {
			for orgName, org := range ordererConfig.Organizations() {
				var certs [][]byte
				certs = append(certs, org.MSP().GetTLSRootCerts()...)
				certs = append(certs, org.MSP().GetTLSIntermediateCerts()...)

				orgAddresses[orgName] = orderers.OrdererOrg{
					Addresses: org.Endpoints(),
					RootCerts: certs,
				}
			}
		}
		ordererSource.Update(globalAddresses, orgAddresses)
	}

	channel := &Channel{
		ledger:         l,
		resources:      bundle,
		cryptoProvider: p.CryptoProvider,
	}

	channel.bundleSource = channelconfig.NewBundleSource(
		bundle,
		ordererSourceCallback,
		gossipCallbackWrapper,
		trustedRootsCallbackWrapper,
		mspCallback,
		channel.bundleUpdate,
	)

	committer := committer.NewLedgerCommitter(l)
	validator := &txvalidator.ValidationRouter{
		CapabilityProvider: channel,
		V14Validator: validatorv14.NewTxValidator(
			cid,
			p.validationWorkersSemaphore,
			channel,
			p.pluginMapper,
			p.CryptoProvider,
		),
		V20Validator: validatorv20.NewTxValidator(
			cid,
			p.validationWorkersSemaphore,
			channel,
			channel.Ledger(),
			&vir.ValidationInfoRetrieveShim{
				New:    newLifecycleValidation,
				Legacy: legacyLifecycleValidation,
			},
			&CollectionInfoShim{
				CollectionAndLifecycleResources: newLifecycleValidation,
				ChannelID:                       bundle.ConfigtxValidator().ChannelID(),
			},
			p.pluginMapper,
			policies.PolicyManagerGetterFunc(p.GetPolicyManager),
			p.CryptoProvider,
		),
	}

	// TODO: does someone need to call Close() on the transientStoreFactory at shutdown of the peer?
	store, err := p.openStore(bundle.ConfigtxValidator().ChannelID())
	if err != nil {
		return errors.Wrapf(err, "[channel %s] failed opening transient store", bundle.ConfigtxValidator().ChannelID())
	}
	channel.store = store

	simpleCollectionStore := privdata.NewSimpleCollectionStore(l, deployedCCInfoProvider)
	p.GossipService.InitializeChannel(bundle.ConfigtxValidator().ChannelID(), ordererSource, store, gossipservice.Support{
		Validator:       validator,
		Committer:       committer,
		CollectionStore: simpleCollectionStore,
		IdDeserializeFactory: gossipprivdata.IdentityDeserializerFactoryFunc(func(chainID string) msp.IdentityDeserializer {
			return mspmgmt.GetManagerForChain(chainID)
		}),
		CapabilityProvider: channel,
	})

	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.channels == nil {
		p.channels = map[string]*Channel{}
	}
	p.channels[cid] = channel

	return nil
}

func (p *Peer) Channel(cid string) *Channel {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	if c, ok := p.channels[cid]; ok {
		return c
	}
	return nil
}

func (p *Peer) StoreForChannel(cid string) *transientstore.Store {
	if c := p.Channel(cid); c != nil {
		return c.Store()
	}
	return nil
}

// GetChannelsInfo returns an array with information about all channels for
// this peer.
func (p *Peer) GetChannelsInfo() []*pb.ChannelInfo {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var channelInfos []*pb.ChannelInfo
	for key := range p.channels {
		ci := &pb.ChannelInfo{ChannelId: key}
		channelInfos = append(channelInfos, ci)
	}
	return channelInfos
}

// GetChannelConfig returns the channel configuration of the channel with channel ID. Note that this
// call returns nil if channel cid has not been created.
func (p *Peer) GetChannelConfig(cid string) channelconfig.Resources {
	if c := p.Channel(cid); c != nil {
		return c.Resources()
	}
	return nil
}

// GetStableChannelConfig returns the stable channel configuration of the channel with channel ID.
// Note that this call returns nil if channel cid has not been created.
func (p *Peer) GetStableChannelConfig(cid string) channelconfig.Resources {
	if c := p.Channel(cid); c != nil {
		return c.Resources()
	}
	return nil
}

// GetLedger returns the ledger of the channel with channel ID. Note that this
// call returns nil if channel cid has not been created.
func (p *Peer) GetLedger(cid string) ledger.PeerLedger {
	if c := p.Channel(cid); c != nil {
		return c.Ledger()
	}
	return nil
}

// GetMSPIDs returns the ID of each application MSP defined on this channel
func (p *Peer) GetMSPIDs(cid string) []string {
	if c := p.Channel(cid); c != nil {
		return c.GetMSPIDs()
	}
	return nil
}

// GetPolicyManager returns the policy manager of the channel with channel ID. Note that this
// call returns nil if channel cid has not been created.
func (p *Peer) GetPolicyManager(cid string) policies.Manager {
	if c := p.Channel(cid); c != nil {
		return c.Resources().PolicyManager()
	}
	return nil
}

// JoinBySnaphotStatus queries ledger mgr to get the status of joinbysnapshot
func (p *Peer) JoinBySnaphotStatus() *pb.JoinBySnapshotStatus {
	return p.LedgerMgr.JoinBySnapshotStatus()
}

// initChannel takes care to initialize channel after peer joined, for example deploys system CCs
func (p *Peer) initChannel(cid string) {
	if p.channelInitializer != nil {
		// Initialize chaincode, namely deploy system CC
		peerLogger.Debugf("Initializing channel %s", cid)
		p.channelInitializer(cid)
	}
}

func (p *Peer) GetApplicationConfig(cid string) (channelconfig.Application, bool) {
	cc := p.GetChannelConfig(cid)
	if cc == nil {
		return nil, false
	}

	return cc.ApplicationConfig()
}

// Initialize sets up any channels that the peer has from the persistence. This
// function should be called at the start up when the ledger and gossip
// ready
func (p *Peer) Initialize(
	init func(string),
	server *comm.GRPCServer,
	pm plugin.Mapper,
	deployedCCInfoProvider ledger.DeployedChaincodeInfoProvider,
	legacyLifecycleValidation plugindispatcher.LifecycleResources,
	newLifecycleValidation plugindispatcher.CollectionAndLifecycleResources,
	nWorkers int,
) {
	// TODO: exported dep fields or constructor
	p.server = server
	p.validationWorkersSemaphore = semaphore.New(nWorkers)
	p.pluginMapper = pm
	p.channelInitializer = init

	ledgerIds, err := p.LedgerMgr.GetLedgerIDs()
	if err != nil {
		panic(fmt.Errorf("error in initializing ledgermgmt: %s", err))
	}

	for _, cid := range ledgerIds {
		peerLogger.Infof("Loading chain %s", cid)
		ledger, err := p.LedgerMgr.OpenLedger(cid)
		if err != nil {
			peerLogger.Errorf("Failed to load ledger %s(%+v)", cid, err)
			peerLogger.Debugf("Error while loading ledger %s with message %s. We continue to the next ledger rather than abort.", cid, err)
			continue
		}
		// Create a chain if we get a valid ledger with config block
		err = p.createChannel(cid, ledger, deployedCCInfoProvider, legacyLifecycleValidation, newLifecycleValidation)
		if err != nil {
			peerLogger.Errorf("Failed to load chain %s(%s)", cid, err)
			peerLogger.Debugf("Error reloading chain %s with message %s. We continue to the next chain rather than abort.", cid, err)
			continue
		}

		p.initChannel(cid)
	}
}
