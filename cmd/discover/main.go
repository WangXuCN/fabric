/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"os"

	"github.com/SmartBFT-Go/fabric/bccsp/factory"
	"github.com/SmartBFT-Go/fabric/cmd/common"
	discovery "github.com/SmartBFT-Go/fabric/discovery/cmd"
)

func main() {
	factory.InitFactories(nil)
	cli := common.NewCLI("discover", "Command line client for fabric discovery service")
	discovery.AddCommands(cli)
	cli.Run(os.Args[1:])
}
