// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/SmartBFT-Go/fabric/common/ledger"
)

type SimpleQueryExecutor struct {
	GetPrivateDataHashStub        func(string, string, string) ([]byte, error)
	getPrivateDataHashMutex       sync.RWMutex
	getPrivateDataHashArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getPrivateDataHashReturns struct {
		result1 []byte
		result2 error
	}
	getPrivateDataHashReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetStateStub        func(string, string) ([]byte, error)
	getStateMutex       sync.RWMutex
	getStateArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getStateReturns struct {
		result1 []byte
		result2 error
	}
	getStateReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetStateRangeScanIteratorStub        func(string, string, string) (ledger.ResultsIterator, error)
	getStateRangeScanIteratorMutex       sync.RWMutex
	getStateRangeScanIteratorArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getStateRangeScanIteratorReturns struct {
		result1 ledger.ResultsIterator
		result2 error
	}
	getStateRangeScanIteratorReturnsOnCall map[int]struct {
		result1 ledger.ResultsIterator
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SimpleQueryExecutor) GetPrivateDataHash(arg1 string, arg2 string, arg3 string) ([]byte, error) {
	fake.getPrivateDataHashMutex.Lock()
	ret, specificReturn := fake.getPrivateDataHashReturnsOnCall[len(fake.getPrivateDataHashArgsForCall)]
	fake.getPrivateDataHashArgsForCall = append(fake.getPrivateDataHashArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetPrivateDataHashStub
	fakeReturns := fake.getPrivateDataHashReturns
	fake.recordInvocation("GetPrivateDataHash", []interface{}{arg1, arg2, arg3})
	fake.getPrivateDataHashMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SimpleQueryExecutor) GetPrivateDataHashCallCount() int {
	fake.getPrivateDataHashMutex.RLock()
	defer fake.getPrivateDataHashMutex.RUnlock()
	return len(fake.getPrivateDataHashArgsForCall)
}

func (fake *SimpleQueryExecutor) GetPrivateDataHashCalls(stub func(string, string, string) ([]byte, error)) {
	fake.getPrivateDataHashMutex.Lock()
	defer fake.getPrivateDataHashMutex.Unlock()
	fake.GetPrivateDataHashStub = stub
}

func (fake *SimpleQueryExecutor) GetPrivateDataHashArgsForCall(i int) (string, string, string) {
	fake.getPrivateDataHashMutex.RLock()
	defer fake.getPrivateDataHashMutex.RUnlock()
	argsForCall := fake.getPrivateDataHashArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SimpleQueryExecutor) GetPrivateDataHashReturns(result1 []byte, result2 error) {
	fake.getPrivateDataHashMutex.Lock()
	defer fake.getPrivateDataHashMutex.Unlock()
	fake.GetPrivateDataHashStub = nil
	fake.getPrivateDataHashReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SimpleQueryExecutor) GetPrivateDataHashReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getPrivateDataHashMutex.Lock()
	defer fake.getPrivateDataHashMutex.Unlock()
	fake.GetPrivateDataHashStub = nil
	if fake.getPrivateDataHashReturnsOnCall == nil {
		fake.getPrivateDataHashReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getPrivateDataHashReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SimpleQueryExecutor) GetState(arg1 string, arg2 string) ([]byte, error) {
	fake.getStateMutex.Lock()
	ret, specificReturn := fake.getStateReturnsOnCall[len(fake.getStateArgsForCall)]
	fake.getStateArgsForCall = append(fake.getStateArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStateStub
	fakeReturns := fake.getStateReturns
	fake.recordInvocation("GetState", []interface{}{arg1, arg2})
	fake.getStateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SimpleQueryExecutor) GetStateCallCount() int {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return len(fake.getStateArgsForCall)
}

func (fake *SimpleQueryExecutor) GetStateCalls(stub func(string, string) ([]byte, error)) {
	fake.getStateMutex.Lock()
	defer fake.getStateMutex.Unlock()
	fake.GetStateStub = stub
}

func (fake *SimpleQueryExecutor) GetStateArgsForCall(i int) (string, string) {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	argsForCall := fake.getStateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SimpleQueryExecutor) GetStateReturns(result1 []byte, result2 error) {
	fake.getStateMutex.Lock()
	defer fake.getStateMutex.Unlock()
	fake.GetStateStub = nil
	fake.getStateReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SimpleQueryExecutor) GetStateReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getStateMutex.Lock()
	defer fake.getStateMutex.Unlock()
	fake.GetStateStub = nil
	if fake.getStateReturnsOnCall == nil {
		fake.getStateReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getStateReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SimpleQueryExecutor) GetStateRangeScanIterator(arg1 string, arg2 string, arg3 string) (ledger.ResultsIterator, error) {
	fake.getStateRangeScanIteratorMutex.Lock()
	ret, specificReturn := fake.getStateRangeScanIteratorReturnsOnCall[len(fake.getStateRangeScanIteratorArgsForCall)]
	fake.getStateRangeScanIteratorArgsForCall = append(fake.getStateRangeScanIteratorArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetStateRangeScanIteratorStub
	fakeReturns := fake.getStateRangeScanIteratorReturns
	fake.recordInvocation("GetStateRangeScanIterator", []interface{}{arg1, arg2, arg3})
	fake.getStateRangeScanIteratorMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SimpleQueryExecutor) GetStateRangeScanIteratorCallCount() int {
	fake.getStateRangeScanIteratorMutex.RLock()
	defer fake.getStateRangeScanIteratorMutex.RUnlock()
	return len(fake.getStateRangeScanIteratorArgsForCall)
}

func (fake *SimpleQueryExecutor) GetStateRangeScanIteratorCalls(stub func(string, string, string) (ledger.ResultsIterator, error)) {
	fake.getStateRangeScanIteratorMutex.Lock()
	defer fake.getStateRangeScanIteratorMutex.Unlock()
	fake.GetStateRangeScanIteratorStub = stub
}

func (fake *SimpleQueryExecutor) GetStateRangeScanIteratorArgsForCall(i int) (string, string, string) {
	fake.getStateRangeScanIteratorMutex.RLock()
	defer fake.getStateRangeScanIteratorMutex.RUnlock()
	argsForCall := fake.getStateRangeScanIteratorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SimpleQueryExecutor) GetStateRangeScanIteratorReturns(result1 ledger.ResultsIterator, result2 error) {
	fake.getStateRangeScanIteratorMutex.Lock()
	defer fake.getStateRangeScanIteratorMutex.Unlock()
	fake.GetStateRangeScanIteratorStub = nil
	fake.getStateRangeScanIteratorReturns = struct {
		result1 ledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *SimpleQueryExecutor) GetStateRangeScanIteratorReturnsOnCall(i int, result1 ledger.ResultsIterator, result2 error) {
	fake.getStateRangeScanIteratorMutex.Lock()
	defer fake.getStateRangeScanIteratorMutex.Unlock()
	fake.GetStateRangeScanIteratorStub = nil
	if fake.getStateRangeScanIteratorReturnsOnCall == nil {
		fake.getStateRangeScanIteratorReturnsOnCall = make(map[int]struct {
			result1 ledger.ResultsIterator
			result2 error
		})
	}
	fake.getStateRangeScanIteratorReturnsOnCall[i] = struct {
		result1 ledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *SimpleQueryExecutor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPrivateDataHashMutex.RLock()
	defer fake.getPrivateDataHashMutex.RUnlock()
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	fake.getStateRangeScanIteratorMutex.RLock()
	defer fake.getStateRangeScanIteratorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SimpleQueryExecutor) recordInvocation(key string, args []interface{}) {
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
