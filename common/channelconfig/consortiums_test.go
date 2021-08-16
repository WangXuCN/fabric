/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package channelconfig

import (
	"testing"

	cb "github.com/SmartBFT-Go/fabric-protos-go/v2/common"
	"github.com/stretchr/testify/require"
)

func TestConsortiums(t *testing.T) {
	_, err := NewConsortiumsConfig(&cb.ConfigGroup{}, nil)
	require.NoError(t, err)
}
