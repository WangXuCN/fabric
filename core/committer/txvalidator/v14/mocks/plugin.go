// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import common "github.com/SmartBFT-Go/fabric-protos-go/v2/common"
import mock "github.com/stretchr/testify/mock"

import validation "github.com/SmartBFT-Go/fabric/core/handlers/validation/api"

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// Init provides a mock function with given fields: dependencies
func (_m *Plugin) Init(dependencies ...validation.Dependency) error {
	_va := make([]interface{}, len(dependencies))
	for _i := range dependencies {
		_va[_i] = dependencies[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...validation.Dependency) error); ok {
		r0 = rf(dependencies...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Validate provides a mock function with given fields: block, namespace, txPosition, actionPosition, contextData
func (_m *Plugin) Validate(block *common.Block, namespace string, txPosition int, actionPosition int, contextData ...validation.ContextDatum) error {
	_va := make([]interface{}, len(contextData))
	for _i := range contextData {
		_va[_i] = contextData[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, block, namespace, txPosition, actionPosition)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Block, string, int, int, ...validation.ContextDatum) error); ok {
		r0 = rf(block, namespace, txPosition, actionPosition, contextData...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
