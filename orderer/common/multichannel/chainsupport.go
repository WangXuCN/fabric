/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package multichannel

import (
	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/common/ledger/blockledger"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/orderer/common/blockcutter"
	"github.com/hyperledger/fabric/orderer/common/msgprocessor"
	"github.com/hyperledger/fabric/orderer/consensus"
	"github.com/hyperledger/fabric/orderer/consensus/smartbft/types"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/orderer"
	"github.com/hyperledger/fabric/protos/orderer/smartbft"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
)

// ChainSupport holds the resources for a particular channel.
type ChainSupport struct {
	*ledgerResources
	msgprocessor.Processor
	*BlockWriter
	consensus.Chain
	cutter blockcutter.Receiver
	crypto.LocalSigner
	CommitteeSizeFetcher func() int
}

func newChainSupport(
	registrar *Registrar,
	ledgerResources *ledgerResources,
	consenters map[string]consensus.Consenter,
	signer crypto.LocalSigner,
	blockcutterMetrics *blockcutter.Metrics,
) *ChainSupport {
	// Read in the last block and metadata for the channel
	lastBlock := blockledger.GetBlock(ledgerResources, ledgerResources.Height()-1)
	metadata, err := utils.GetMetadataFromBlock(lastBlock, cb.BlockMetadataIndex_ORDERER)
	// Assuming a block created with cb.NewBlock(), this should not
	// error even if the orderer metadata is an empty byte slice
	if err != nil {
		logger.Fatalf("[channel: %s] Error extracting orderer metadata: %s", ledgerResources.ConfigtxValidator().ChainID(), err)
	}

	// Construct limited support needed as a parameter for additional support
	cs := &ChainSupport{
		ledgerResources: ledgerResources,
		LocalSigner:     signer,
		cutter: blockcutter.NewReceiverImpl(
			ledgerResources.ConfigtxValidator().ChainID(),
			ledgerResources,
			blockcutterMetrics,
		),
	}

	var synchronousBlockWriting bool
	oc, _ := ledgerResources.OrdererConfig()
	if oc.ConsensusType() == "smartbft" {
		synchronousBlockWriting = true
	}

	// Set up the msgprocessor
	cs.Processor = msgprocessor.NewStandardChannel(cs, msgprocessor.CreateStandardChannelFilters(cs, registrar.config))

	// Set up the block writer
	cs.BlockWriter = newBlockWriter(lastBlock, registrar, cs, synchronousBlockWriting)

	// Set up the consenter
	consenterType := ledgerResources.SharedConfig().ConsensusType()
	consenter, ok := consenters[consenterType]
	if !ok {
		logger.Panicf("Error retrieving consenter of type: %s", consenterType)
	}

	cs.Chain, err = consenter.HandleChain(cs, metadata)
	if err != nil {
		logger.Panicf("[channel: %s] Error creating consenter: %s", cs.ChainID(), err)
	}

	logger.Debugf("[channel: %s] Done creating channel support resources", cs.ChainID())

	return cs
}

// Block returns a block with the following number,
// or nil if such a block doesn't exist.
func (cs *ChainSupport) Block(number uint64) *cb.Block {
	if cs.Height() <= number {
		return nil
	}
	return blockledger.GetBlock(cs.Reader(), number)
}

func (cs *ChainSupport) Reader() blockledger.Reader {
	return cs
}

// Signer returns the crypto.Localsigner for this channel.
func (cs *ChainSupport) Signer() crypto.LocalSigner {
	return cs
}

func (cs *ChainSupport) start() {
	cs.Chain.Start()
}

// BlockCutter returns the blockcutter.Receiver instance for this channel.
func (cs *ChainSupport) BlockCutter() blockcutter.Receiver {
	return cs.cutter
}

// Validate passes through to the underlying configtx.Validator
func (cs *ChainSupport) Validate(configEnv *cb.ConfigEnvelope) error {
	return cs.ConfigtxValidator().Validate(configEnv)
}

// ProposeConfigUpdate passes through to the underlying configtx.Validator
func (cs *ChainSupport) ProposeConfigUpdate(configtx *cb.Envelope) (*cb.ConfigEnvelope, error) {
	env, err := cs.ConfigtxValidator().ProposeConfigUpdate(configtx)
	if err != nil {
		return nil, err
	}

	bundle, err := cs.CreateBundle(cs.ChainID(), env.Config)
	if err != nil {
		return nil, err
	}

	if err = checkResources(bundle); err != nil {
		return nil, errors.Wrap(err, "config update is not compatible")
	}

	return env, cs.ValidateNew(bundle)
}

// ChainID passes through to the underlying configtx.Validator
func (cs *ChainSupport) ChainID() string {
	return cs.ConfigtxValidator().ChainID()
}

// ConfigProto passes through to the underlying configtx.Validator
func (cs *ChainSupport) ConfigProto() *cb.Config {
	return cs.ConfigtxValidator().ConfigProto()
}

// Sequence passes through to the underlying configtx.Validator
func (cs *ChainSupport) Sequence() uint64 {
	return cs.ConfigtxValidator().Sequence()
}

// Append appends a new block to the ledger in its raw form,
// unlike WriteBlock that also mutates its metadata.
func (cs *ChainSupport) Append(block *cb.Block) error {
	return cs.ledgerResources.ReadWriter.Append(block)
}

func (cs *ChainSupport) NodeCountForBlock(seq uint64) int {
	if seq <= 1 {
		return 0
	}

	block := blockledger.GetBlock(cs.support, seq-1)
	if md, _ := types.CommitteeMetadataFromBlock(block); md != nil {
		return int(md.CommitteeSize)
	}
	return 0
}

// VerifyBlockSignature verifies a signature of a block.
// It has an optional argument of a configuration envelope
// which would make the block verification to use validation rules
// based on the given configuration in the ConfigEnvelope.
// If the config envelope passed is nil, then the validation rules used
// are the ones that were applied at commit of previous blocks.
func (cs *ChainSupport) VerifyBlockSignature(sd []*cb.SignedData, envelope *cb.ConfigEnvelope, committeeSize int) error {
	policyMgr := cs.PolicyManager()
	// If the envelope passed isn't nil, we should use a different policy manager.
	if envelope != nil {
		bundle, err := channelconfig.NewBundle(cs.ChainID(), envelope.Config)
		if err != nil {
			return err
		}
		policyMgr = bundle.PolicyManager()
	}
	policy, exists := policyMgr.GetPolicy(policies.BlockValidation)
	if !exists {
		return errors.Errorf("policy %s wasn't found", policies.BlockValidation)
	}
	err := policy.(policies.BFTPolicy).BFTEvaluate(sd, committeeSize)
	if err != nil {
		return errors.Wrap(err, "block verification failed")
	}
	return nil
}

func (cs *ChainSupport) Id2Identity(envelope *cb.ConfigEnvelope) map[uint64][]byte {
	consensusMD := cs.SharedConfig().ConsensusMetadata()
	if envelope != nil {
		consensusMD = envelope.Config.ChannelGroup.Groups[channelconfig.OrdererGroupKey].Values[channelconfig.ConsensusTypeKey].Value
		ct := &orderer.ConsensusType{}
		err := proto.Unmarshal(consensusMD, ct)
		if err != nil {
			logger.Panicf("Failed unmarshaling ConsensusType from consensusType: %v", err)
		}
		consensusMD = ct.Metadata
	}

	m := &smartbft.ConfigMetadata{}
	err := proto.Unmarshal(consensusMD, m)
	if err != nil {
		logger.Panicf("Failed unmarshaling ConfigMetadata from metadata: %v", err)
	}

	res := make(map[uint64][]byte)
	for _, consenter := range m.Consenters {
		res[consenter.ConsenterId] = consenter.Identity
	}

	return res
}
