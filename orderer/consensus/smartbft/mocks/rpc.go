// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import orderer "github.com/SmartBFT-Go/fabric-protos-go/v2/orderer"

// RPC is an autogenerated mock type for the RPC type
type RPC struct {
	mock.Mock
}

// SendConsensus provides a mock function with given fields: dest, msg
func (_m *RPC) SendConsensus(dest uint64, msg *orderer.ConsensusRequest) error {
	ret := _m.Called(dest, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, *orderer.ConsensusRequest) error); ok {
		r0 = rf(dest, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendSubmit provides a mock function with given fields: dest, request
func (_m *RPC) SendSubmit(dest uint64, request *orderer.SubmitRequest) error {
	ret := _m.Called(dest, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, *orderer.SubmitRequest) error); ok {
		r0 = rf(dest, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
