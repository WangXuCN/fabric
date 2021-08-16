// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import channelconfig "github.com/SmartBFT-Go/fabric/common/channelconfig"
import common "github.com/SmartBFT-Go/fabric-protos-go/v2/common"
import mock "github.com/stretchr/testify/mock"

// ChannelConfigTemplator is an autogenerated mock type for the ChannelConfigTemplator type
type ChannelConfigTemplator struct {
	mock.Mock
}

// NewChannelConfig provides a mock function with given fields: env
func (_m *ChannelConfigTemplator) NewChannelConfig(env *common.Envelope) (channelconfig.Resources, error) {
	ret := _m.Called(env)

	var r0 channelconfig.Resources
	if rf, ok := ret.Get(0).(func(*common.Envelope) channelconfig.Resources); ok {
		r0 = rf(env)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(channelconfig.Resources)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Envelope) error); ok {
		r1 = rf(env)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
