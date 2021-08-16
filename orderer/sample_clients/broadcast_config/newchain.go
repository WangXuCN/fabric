// Copyright IBM Corp. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	cb "github.com/SmartBFT-Go/fabric-protos-go/v2/common"
	"github.com/SmartBFT-Go/fabric/internal/configtxgen/encoder"
	"github.com/SmartBFT-Go/fabric/internal/configtxgen/genesisconfig"
	"github.com/SmartBFT-Go/fabric/internal/pkg/identity"
)

func newChainRequest(
	consensusType,
	creationPolicy,
	newChannelID string,
	signer identity.SignerSerializer,
) *cb.Envelope {
	env, err := encoder.MakeChannelCreationTransaction(
		newChannelID,
		signer,
		genesisconfig.Load(genesisconfig.SampleSingleMSPChannelProfile),
	)
	if err != nil {
		panic(err)
	}
	return env
}
