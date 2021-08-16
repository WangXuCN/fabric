// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/SmartBFT-Go/fabric/core/ledger"
)

type ChaincodeLifecycleEventProvider struct {
	RegisterListenerStub        func(string, ledger.ChaincodeLifecycleEventListener, bool) error
	registerListenerMutex       sync.RWMutex
	registerListenerArgsForCall []struct {
		arg1 string
		arg2 ledger.ChaincodeLifecycleEventListener
		arg3 bool
	}
	registerListenerReturns struct {
		result1 error
	}
	registerListenerReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChaincodeLifecycleEventProvider) RegisterListener(arg1 string, arg2 ledger.ChaincodeLifecycleEventListener, arg3 bool) error {
	fake.registerListenerMutex.Lock()
	ret, specificReturn := fake.registerListenerReturnsOnCall[len(fake.registerListenerArgsForCall)]
	fake.registerListenerArgsForCall = append(fake.registerListenerArgsForCall, struct {
		arg1 string
		arg2 ledger.ChaincodeLifecycleEventListener
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.RegisterListenerStub
	fakeReturns := fake.registerListenerReturns
	fake.recordInvocation("RegisterListener", []interface{}{arg1, arg2, arg3})
	fake.registerListenerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ChaincodeLifecycleEventProvider) RegisterListenerCallCount() int {
	fake.registerListenerMutex.RLock()
	defer fake.registerListenerMutex.RUnlock()
	return len(fake.registerListenerArgsForCall)
}

func (fake *ChaincodeLifecycleEventProvider) RegisterListenerCalls(stub func(string, ledger.ChaincodeLifecycleEventListener, bool) error) {
	fake.registerListenerMutex.Lock()
	defer fake.registerListenerMutex.Unlock()
	fake.RegisterListenerStub = stub
}

func (fake *ChaincodeLifecycleEventProvider) RegisterListenerArgsForCall(i int) (string, ledger.ChaincodeLifecycleEventListener, bool) {
	fake.registerListenerMutex.RLock()
	defer fake.registerListenerMutex.RUnlock()
	argsForCall := fake.registerListenerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChaincodeLifecycleEventProvider) RegisterListenerReturns(result1 error) {
	fake.registerListenerMutex.Lock()
	defer fake.registerListenerMutex.Unlock()
	fake.RegisterListenerStub = nil
	fake.registerListenerReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeLifecycleEventProvider) RegisterListenerReturnsOnCall(i int, result1 error) {
	fake.registerListenerMutex.Lock()
	defer fake.registerListenerMutex.Unlock()
	fake.RegisterListenerStub = nil
	if fake.registerListenerReturnsOnCall == nil {
		fake.registerListenerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerListenerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeLifecycleEventProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerListenerMutex.RLock()
	defer fake.registerListenerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChaincodeLifecycleEventProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ ledger.ChaincodeLifecycleEventProvider = new(ChaincodeLifecycleEventProvider)
