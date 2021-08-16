/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"github.com/SmartBFT-Go/fabric-chaincode-go/shim"
	"github.com/SmartBFT-Go/fabric/integration/chaincode/marbles_private"
)

func main() {
	err := shim.Start(&marbles_private.MarblesPrivateChaincode{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting Simple chaincode: %s", err)
		os.Exit(2)
	}
}
