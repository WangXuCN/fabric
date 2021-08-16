// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import discovery "github.com/SmartBFT-Go/fabric/discovery/cmd"
import mock "github.com/stretchr/testify/mock"

// ResponseParser is an autogenerated mock type for the ResponseParser type
type ResponseParser struct {
	mock.Mock
}

// ParseResponse provides a mock function with given fields: channel, response
func (_m *ResponseParser) ParseResponse(channel string, response discovery.ServiceResponse) error {
	ret := _m.Called(channel, response)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, discovery.ServiceResponse) error); ok {
		r0 = rf(channel, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
