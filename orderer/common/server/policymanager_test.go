/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package server

import (
	"testing"

	"github.com/SmartBFT-Go/fabric/bccsp/sw"
	"github.com/SmartBFT-Go/fabric/common/channelconfig"
	"github.com/SmartBFT-Go/fabric/common/flogging"
	"github.com/SmartBFT-Go/fabric/common/genesis"
	"github.com/SmartBFT-Go/fabric/common/policies"
	"github.com/SmartBFT-Go/fabric/core/config/configtest"
	"github.com/SmartBFT-Go/fabric/internal/configtxgen/encoder"
	"github.com/SmartBFT-Go/fabric/internal/configtxgen/genesisconfig"
	"github.com/SmartBFT-Go/fabric/protoutil"
	"github.com/stretchr/testify/assert"
)

func TestDynamicPolicyManager(t *testing.T) {
	profile := genesisconfig.Load(genesisconfig.SampleDevModeSoloProfile, configtest.GetDevConfigDir())
	channelGroup, err := encoder.NewChannelGroup(profile)
	assert.NoError(t, err)

	block := genesis.NewFactoryImpl(channelGroup).Block("test")
	env := protoutil.UnmarshalEnvelopeOrPanic(block.Data.Data[0])
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)

	bundle, err := channelconfig.NewBundleFromEnvelope(env, cryptoProvider)
	assert.NoError(t, err)

	l := flogging.MustGetLogger("test")
	dpmr := &DynamicPolicyManagerRegistry{
		Logger: l,
	}

	dpmr.Update(bundle)
	managerByChain := dpmr.Registry()

	for _, testCase := range []struct {
		description string
		channel     string
		succeeds    bool
	}{
		{
			description: "succeeds",
			channel:     "test",
			succeeds:    true,
		},
		{
			channel:     "not test",
			description: "fails",
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			testManager := managerByChain(testCase.channel)
			assert.NotNil(t, testManager)

			pol, ok := testManager.GetPolicy(policies.ChannelReaders)
			assert.Equal(t, testCase.succeeds, ok)
			if testCase.succeeds {
				assert.NotNil(t, pol)
			} else {
				assert.Nil(t, pol)
			}

			mgr, ok := testManager.Manager([]string{policies.OrdererPrefix})
			assert.Equal(t, testCase.succeeds, ok)
			if testCase.succeeds {
				assert.NotNil(t, mgr)
			} else {
				assert.Nil(t, mgr)
			}
		})
	}
}
