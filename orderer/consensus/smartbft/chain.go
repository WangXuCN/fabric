/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package smartbft

import "github.com/hyperledger/fabric/protos/common"

type bftchain struct{}

func (c *bftchain) Order(env *common.Envelope, configSeq uint64) error {
	// TODO: Implement Order, add implementation to order normal transactions
	return nil
}

func (c *bftchain) Configure(config *common.Envelope, configSeq uint64) error {
	// TODO: Implement Configure, add implementation to process configuration updates
	return nil
}

func (c *bftchain) WaitReady() error {
	// TODO: Implement WaitReady
	return nil
}

func (c *bftchain) Errored() <-chan struct{} {
	// TODO: Implement Errored
	return nil
}

func (c *bftchain) Start() {
	// TODO: Implement Start
}

func (c *bftchain) Halt() {
	// TODO: Implement Halt
}
