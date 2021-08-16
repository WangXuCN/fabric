// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/SmartBFT-Go/fabric-protos-go/v2/common"
	"github.com/SmartBFT-Go/fabric-protos-go/v2/orderer"
	"github.com/hyperledger/fabric/common/ledger/blockledger"
)

type ReadWriter struct {
	AppendStub        func(*common.Block) error
	appendMutex       sync.RWMutex
	appendArgsForCall []struct {
		arg1 *common.Block
	}
	appendReturns struct {
		result1 error
	}
	appendReturnsOnCall map[int]struct {
		result1 error
	}
	HeightStub        func() uint64
	heightMutex       sync.RWMutex
	heightArgsForCall []struct {
	}
	heightReturns struct {
		result1 uint64
	}
	heightReturnsOnCall map[int]struct {
		result1 uint64
	}
	IteratorStub        func(*orderer.SeekPosition) (blockledger.Iterator, uint64)
	iteratorMutex       sync.RWMutex
	iteratorArgsForCall []struct {
		arg1 *orderer.SeekPosition
	}
	iteratorReturns struct {
		result1 blockledger.Iterator
		result2 uint64
	}
	iteratorReturnsOnCall map[int]struct {
		result1 blockledger.Iterator
		result2 uint64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReadWriter) Append(arg1 *common.Block) error {
	fake.appendMutex.Lock()
	ret, specificReturn := fake.appendReturnsOnCall[len(fake.appendArgsForCall)]
	fake.appendArgsForCall = append(fake.appendArgsForCall, struct {
		arg1 *common.Block
	}{arg1})
	stub := fake.AppendStub
	fakeReturns := fake.appendReturns
	fake.recordInvocation("Append", []interface{}{arg1})
	fake.appendMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ReadWriter) AppendCallCount() int {
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	return len(fake.appendArgsForCall)
}

func (fake *ReadWriter) AppendCalls(stub func(*common.Block) error) {
	fake.appendMutex.Lock()
	defer fake.appendMutex.Unlock()
	fake.AppendStub = stub
}

func (fake *ReadWriter) AppendArgsForCall(i int) *common.Block {
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	argsForCall := fake.appendArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ReadWriter) AppendReturns(result1 error) {
	fake.appendMutex.Lock()
	defer fake.appendMutex.Unlock()
	fake.AppendStub = nil
	fake.appendReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReadWriter) AppendReturnsOnCall(i int, result1 error) {
	fake.appendMutex.Lock()
	defer fake.appendMutex.Unlock()
	fake.AppendStub = nil
	if fake.appendReturnsOnCall == nil {
		fake.appendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.appendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ReadWriter) Height() uint64 {
	fake.heightMutex.Lock()
	ret, specificReturn := fake.heightReturnsOnCall[len(fake.heightArgsForCall)]
	fake.heightArgsForCall = append(fake.heightArgsForCall, struct {
	}{})
	stub := fake.HeightStub
	fakeReturns := fake.heightReturns
	fake.recordInvocation("Height", []interface{}{})
	fake.heightMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ReadWriter) HeightCallCount() int {
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	return len(fake.heightArgsForCall)
}

func (fake *ReadWriter) HeightCalls(stub func() uint64) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = stub
}

func (fake *ReadWriter) HeightReturns(result1 uint64) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = nil
	fake.heightReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *ReadWriter) HeightReturnsOnCall(i int, result1 uint64) {
	fake.heightMutex.Lock()
	defer fake.heightMutex.Unlock()
	fake.HeightStub = nil
	if fake.heightReturnsOnCall == nil {
		fake.heightReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.heightReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *ReadWriter) Iterator(arg1 *orderer.SeekPosition) (blockledger.Iterator, uint64) {
	fake.iteratorMutex.Lock()
	ret, specificReturn := fake.iteratorReturnsOnCall[len(fake.iteratorArgsForCall)]
	fake.iteratorArgsForCall = append(fake.iteratorArgsForCall, struct {
		arg1 *orderer.SeekPosition
	}{arg1})
	stub := fake.IteratorStub
	fakeReturns := fake.iteratorReturns
	fake.recordInvocation("Iterator", []interface{}{arg1})
	fake.iteratorMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReadWriter) IteratorCallCount() int {
	fake.iteratorMutex.RLock()
	defer fake.iteratorMutex.RUnlock()
	return len(fake.iteratorArgsForCall)
}

func (fake *ReadWriter) IteratorCalls(stub func(*orderer.SeekPosition) (blockledger.Iterator, uint64)) {
	fake.iteratorMutex.Lock()
	defer fake.iteratorMutex.Unlock()
	fake.IteratorStub = stub
}

func (fake *ReadWriter) IteratorArgsForCall(i int) *orderer.SeekPosition {
	fake.iteratorMutex.RLock()
	defer fake.iteratorMutex.RUnlock()
	argsForCall := fake.iteratorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ReadWriter) IteratorReturns(result1 blockledger.Iterator, result2 uint64) {
	fake.iteratorMutex.Lock()
	defer fake.iteratorMutex.Unlock()
	fake.IteratorStub = nil
	fake.iteratorReturns = struct {
		result1 blockledger.Iterator
		result2 uint64
	}{result1, result2}
}

func (fake *ReadWriter) IteratorReturnsOnCall(i int, result1 blockledger.Iterator, result2 uint64) {
	fake.iteratorMutex.Lock()
	defer fake.iteratorMutex.Unlock()
	fake.IteratorStub = nil
	if fake.iteratorReturnsOnCall == nil {
		fake.iteratorReturnsOnCall = make(map[int]struct {
			result1 blockledger.Iterator
			result2 uint64
		})
	}
	fake.iteratorReturnsOnCall[i] = struct {
		result1 blockledger.Iterator
		result2 uint64
	}{result1, result2}
}

func (fake *ReadWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	fake.heightMutex.RLock()
	defer fake.heightMutex.RUnlock()
	fake.iteratorMutex.RLock()
	defer fake.iteratorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReadWriter) recordInvocation(key string, args []interface{}) {
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
