// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"time"

	mspa "github.com/SmartBFT-Go/fabric-protos-go/v2/msp"
	"github.com/SmartBFT-Go/fabric/msp"
)

type SigningIdentity struct {
	AnonymousStub        func() bool
	anonymousMutex       sync.RWMutex
	anonymousArgsForCall []struct {
	}
	anonymousReturns struct {
		result1 bool
	}
	anonymousReturnsOnCall map[int]struct {
		result1 bool
	}
	ExpiresAtStub        func() time.Time
	expiresAtMutex       sync.RWMutex
	expiresAtArgsForCall []struct {
	}
	expiresAtReturns struct {
		result1 time.Time
	}
	expiresAtReturnsOnCall map[int]struct {
		result1 time.Time
	}
	GetIdentifierStub        func() *msp.IdentityIdentifier
	getIdentifierMutex       sync.RWMutex
	getIdentifierArgsForCall []struct {
	}
	getIdentifierReturns struct {
		result1 *msp.IdentityIdentifier
	}
	getIdentifierReturnsOnCall map[int]struct {
		result1 *msp.IdentityIdentifier
	}
	GetMSPIdentifierStub        func() string
	getMSPIdentifierMutex       sync.RWMutex
	getMSPIdentifierArgsForCall []struct {
	}
	getMSPIdentifierReturns struct {
		result1 string
	}
	getMSPIdentifierReturnsOnCall map[int]struct {
		result1 string
	}
	GetOrganizationalUnitsStub        func() []*msp.OUIdentifier
	getOrganizationalUnitsMutex       sync.RWMutex
	getOrganizationalUnitsArgsForCall []struct {
	}
	getOrganizationalUnitsReturns struct {
		result1 []*msp.OUIdentifier
	}
	getOrganizationalUnitsReturnsOnCall map[int]struct {
		result1 []*msp.OUIdentifier
	}
	GetPublicVersionStub        func() msp.Identity
	getPublicVersionMutex       sync.RWMutex
	getPublicVersionArgsForCall []struct {
	}
	getPublicVersionReturns struct {
		result1 msp.Identity
	}
	getPublicVersionReturnsOnCall map[int]struct {
		result1 msp.Identity
	}
	SatisfiesPrincipalStub        func(*mspa.MSPPrincipal) error
	satisfiesPrincipalMutex       sync.RWMutex
	satisfiesPrincipalArgsForCall []struct {
		arg1 *mspa.MSPPrincipal
	}
	satisfiesPrincipalReturns struct {
		result1 error
	}
	satisfiesPrincipalReturnsOnCall map[int]struct {
		result1 error
	}
	SerializeStub        func() ([]byte, error)
	serializeMutex       sync.RWMutex
	serializeArgsForCall []struct {
	}
	serializeReturns struct {
		result1 []byte
		result2 error
	}
	serializeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	SignStub        func([]byte) ([]byte, error)
	signMutex       sync.RWMutex
	signArgsForCall []struct {
		arg1 []byte
	}
	signReturns struct {
		result1 []byte
		result2 error
	}
	signReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	ValidateStub        func() error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
	}
	validateReturns struct {
		result1 error
	}
	validateReturnsOnCall map[int]struct {
		result1 error
	}
	VerifyStub        func([]byte, []byte) error
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		arg1 []byte
		arg2 []byte
	}
	verifyReturns struct {
		result1 error
	}
	verifyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SigningIdentity) Anonymous() bool {
	fake.anonymousMutex.Lock()
	ret, specificReturn := fake.anonymousReturnsOnCall[len(fake.anonymousArgsForCall)]
	fake.anonymousArgsForCall = append(fake.anonymousArgsForCall, struct {
	}{})
	stub := fake.AnonymousStub
	fakeReturns := fake.anonymousReturns
	fake.recordInvocation("Anonymous", []interface{}{})
	fake.anonymousMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SigningIdentity) AnonymousCallCount() int {
	fake.anonymousMutex.RLock()
	defer fake.anonymousMutex.RUnlock()
	return len(fake.anonymousArgsForCall)
}

func (fake *SigningIdentity) AnonymousCalls(stub func() bool) {
	fake.anonymousMutex.Lock()
	defer fake.anonymousMutex.Unlock()
	fake.AnonymousStub = stub
}

func (fake *SigningIdentity) AnonymousReturns(result1 bool) {
	fake.anonymousMutex.Lock()
	defer fake.anonymousMutex.Unlock()
	fake.AnonymousStub = nil
	fake.anonymousReturns = struct {
		result1 bool
	}{result1}
}

func (fake *SigningIdentity) AnonymousReturnsOnCall(i int, result1 bool) {
	fake.anonymousMutex.Lock()
	defer fake.anonymousMutex.Unlock()
	fake.AnonymousStub = nil
	if fake.anonymousReturnsOnCall == nil {
		fake.anonymousReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.anonymousReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *SigningIdentity) ExpiresAt() time.Time {
	fake.expiresAtMutex.Lock()
	ret, specificReturn := fake.expiresAtReturnsOnCall[len(fake.expiresAtArgsForCall)]
	fake.expiresAtArgsForCall = append(fake.expiresAtArgsForCall, struct {
	}{})
	stub := fake.ExpiresAtStub
	fakeReturns := fake.expiresAtReturns
	fake.recordInvocation("ExpiresAt", []interface{}{})
	fake.expiresAtMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SigningIdentity) ExpiresAtCallCount() int {
	fake.expiresAtMutex.RLock()
	defer fake.expiresAtMutex.RUnlock()
	return len(fake.expiresAtArgsForCall)
}

func (fake *SigningIdentity) ExpiresAtCalls(stub func() time.Time) {
	fake.expiresAtMutex.Lock()
	defer fake.expiresAtMutex.Unlock()
	fake.ExpiresAtStub = stub
}

func (fake *SigningIdentity) ExpiresAtReturns(result1 time.Time) {
	fake.expiresAtMutex.Lock()
	defer fake.expiresAtMutex.Unlock()
	fake.ExpiresAtStub = nil
	fake.expiresAtReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *SigningIdentity) ExpiresAtReturnsOnCall(i int, result1 time.Time) {
	fake.expiresAtMutex.Lock()
	defer fake.expiresAtMutex.Unlock()
	fake.ExpiresAtStub = nil
	if fake.expiresAtReturnsOnCall == nil {
		fake.expiresAtReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.expiresAtReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *SigningIdentity) GetIdentifier() *msp.IdentityIdentifier {
	fake.getIdentifierMutex.Lock()
	ret, specificReturn := fake.getIdentifierReturnsOnCall[len(fake.getIdentifierArgsForCall)]
	fake.getIdentifierArgsForCall = append(fake.getIdentifierArgsForCall, struct {
	}{})
	stub := fake.GetIdentifierStub
	fakeReturns := fake.getIdentifierReturns
	fake.recordInvocation("GetIdentifier", []interface{}{})
	fake.getIdentifierMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SigningIdentity) GetIdentifierCallCount() int {
	fake.getIdentifierMutex.RLock()
	defer fake.getIdentifierMutex.RUnlock()
	return len(fake.getIdentifierArgsForCall)
}

func (fake *SigningIdentity) GetIdentifierCalls(stub func() *msp.IdentityIdentifier) {
	fake.getIdentifierMutex.Lock()
	defer fake.getIdentifierMutex.Unlock()
	fake.GetIdentifierStub = stub
}

func (fake *SigningIdentity) GetIdentifierReturns(result1 *msp.IdentityIdentifier) {
	fake.getIdentifierMutex.Lock()
	defer fake.getIdentifierMutex.Unlock()
	fake.GetIdentifierStub = nil
	fake.getIdentifierReturns = struct {
		result1 *msp.IdentityIdentifier
	}{result1}
}

func (fake *SigningIdentity) GetIdentifierReturnsOnCall(i int, result1 *msp.IdentityIdentifier) {
	fake.getIdentifierMutex.Lock()
	defer fake.getIdentifierMutex.Unlock()
	fake.GetIdentifierStub = nil
	if fake.getIdentifierReturnsOnCall == nil {
		fake.getIdentifierReturnsOnCall = make(map[int]struct {
			result1 *msp.IdentityIdentifier
		})
	}
	fake.getIdentifierReturnsOnCall[i] = struct {
		result1 *msp.IdentityIdentifier
	}{result1}
}

func (fake *SigningIdentity) GetMSPIdentifier() string {
	fake.getMSPIdentifierMutex.Lock()
	ret, specificReturn := fake.getMSPIdentifierReturnsOnCall[len(fake.getMSPIdentifierArgsForCall)]
	fake.getMSPIdentifierArgsForCall = append(fake.getMSPIdentifierArgsForCall, struct {
	}{})
	stub := fake.GetMSPIdentifierStub
	fakeReturns := fake.getMSPIdentifierReturns
	fake.recordInvocation("GetMSPIdentifier", []interface{}{})
	fake.getMSPIdentifierMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SigningIdentity) GetMSPIdentifierCallCount() int {
	fake.getMSPIdentifierMutex.RLock()
	defer fake.getMSPIdentifierMutex.RUnlock()
	return len(fake.getMSPIdentifierArgsForCall)
}

func (fake *SigningIdentity) GetMSPIdentifierCalls(stub func() string) {
	fake.getMSPIdentifierMutex.Lock()
	defer fake.getMSPIdentifierMutex.Unlock()
	fake.GetMSPIdentifierStub = stub
}

func (fake *SigningIdentity) GetMSPIdentifierReturns(result1 string) {
	fake.getMSPIdentifierMutex.Lock()
	defer fake.getMSPIdentifierMutex.Unlock()
	fake.GetMSPIdentifierStub = nil
	fake.getMSPIdentifierReturns = struct {
		result1 string
	}{result1}
}

func (fake *SigningIdentity) GetMSPIdentifierReturnsOnCall(i int, result1 string) {
	fake.getMSPIdentifierMutex.Lock()
	defer fake.getMSPIdentifierMutex.Unlock()
	fake.GetMSPIdentifierStub = nil
	if fake.getMSPIdentifierReturnsOnCall == nil {
		fake.getMSPIdentifierReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getMSPIdentifierReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *SigningIdentity) GetOrganizationalUnits() []*msp.OUIdentifier {
	fake.getOrganizationalUnitsMutex.Lock()
	ret, specificReturn := fake.getOrganizationalUnitsReturnsOnCall[len(fake.getOrganizationalUnitsArgsForCall)]
	fake.getOrganizationalUnitsArgsForCall = append(fake.getOrganizationalUnitsArgsForCall, struct {
	}{})
	stub := fake.GetOrganizationalUnitsStub
	fakeReturns := fake.getOrganizationalUnitsReturns
	fake.recordInvocation("GetOrganizationalUnits", []interface{}{})
	fake.getOrganizationalUnitsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SigningIdentity) GetOrganizationalUnitsCallCount() int {
	fake.getOrganizationalUnitsMutex.RLock()
	defer fake.getOrganizationalUnitsMutex.RUnlock()
	return len(fake.getOrganizationalUnitsArgsForCall)
}

func (fake *SigningIdentity) GetOrganizationalUnitsCalls(stub func() []*msp.OUIdentifier) {
	fake.getOrganizationalUnitsMutex.Lock()
	defer fake.getOrganizationalUnitsMutex.Unlock()
	fake.GetOrganizationalUnitsStub = stub
}

func (fake *SigningIdentity) GetOrganizationalUnitsReturns(result1 []*msp.OUIdentifier) {
	fake.getOrganizationalUnitsMutex.Lock()
	defer fake.getOrganizationalUnitsMutex.Unlock()
	fake.GetOrganizationalUnitsStub = nil
	fake.getOrganizationalUnitsReturns = struct {
		result1 []*msp.OUIdentifier
	}{result1}
}

func (fake *SigningIdentity) GetOrganizationalUnitsReturnsOnCall(i int, result1 []*msp.OUIdentifier) {
	fake.getOrganizationalUnitsMutex.Lock()
	defer fake.getOrganizationalUnitsMutex.Unlock()
	fake.GetOrganizationalUnitsStub = nil
	if fake.getOrganizationalUnitsReturnsOnCall == nil {
		fake.getOrganizationalUnitsReturnsOnCall = make(map[int]struct {
			result1 []*msp.OUIdentifier
		})
	}
	fake.getOrganizationalUnitsReturnsOnCall[i] = struct {
		result1 []*msp.OUIdentifier
	}{result1}
}

func (fake *SigningIdentity) GetPublicVersion() msp.Identity {
	fake.getPublicVersionMutex.Lock()
	ret, specificReturn := fake.getPublicVersionReturnsOnCall[len(fake.getPublicVersionArgsForCall)]
	fake.getPublicVersionArgsForCall = append(fake.getPublicVersionArgsForCall, struct {
	}{})
	stub := fake.GetPublicVersionStub
	fakeReturns := fake.getPublicVersionReturns
	fake.recordInvocation("GetPublicVersion", []interface{}{})
	fake.getPublicVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SigningIdentity) GetPublicVersionCallCount() int {
	fake.getPublicVersionMutex.RLock()
	defer fake.getPublicVersionMutex.RUnlock()
	return len(fake.getPublicVersionArgsForCall)
}

func (fake *SigningIdentity) GetPublicVersionCalls(stub func() msp.Identity) {
	fake.getPublicVersionMutex.Lock()
	defer fake.getPublicVersionMutex.Unlock()
	fake.GetPublicVersionStub = stub
}

func (fake *SigningIdentity) GetPublicVersionReturns(result1 msp.Identity) {
	fake.getPublicVersionMutex.Lock()
	defer fake.getPublicVersionMutex.Unlock()
	fake.GetPublicVersionStub = nil
	fake.getPublicVersionReturns = struct {
		result1 msp.Identity
	}{result1}
}

func (fake *SigningIdentity) GetPublicVersionReturnsOnCall(i int, result1 msp.Identity) {
	fake.getPublicVersionMutex.Lock()
	defer fake.getPublicVersionMutex.Unlock()
	fake.GetPublicVersionStub = nil
	if fake.getPublicVersionReturnsOnCall == nil {
		fake.getPublicVersionReturnsOnCall = make(map[int]struct {
			result1 msp.Identity
		})
	}
	fake.getPublicVersionReturnsOnCall[i] = struct {
		result1 msp.Identity
	}{result1}
}

func (fake *SigningIdentity) SatisfiesPrincipal(arg1 *mspa.MSPPrincipal) error {
	fake.satisfiesPrincipalMutex.Lock()
	ret, specificReturn := fake.satisfiesPrincipalReturnsOnCall[len(fake.satisfiesPrincipalArgsForCall)]
	fake.satisfiesPrincipalArgsForCall = append(fake.satisfiesPrincipalArgsForCall, struct {
		arg1 *mspa.MSPPrincipal
	}{arg1})
	stub := fake.SatisfiesPrincipalStub
	fakeReturns := fake.satisfiesPrincipalReturns
	fake.recordInvocation("SatisfiesPrincipal", []interface{}{arg1})
	fake.satisfiesPrincipalMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SigningIdentity) SatisfiesPrincipalCallCount() int {
	fake.satisfiesPrincipalMutex.RLock()
	defer fake.satisfiesPrincipalMutex.RUnlock()
	return len(fake.satisfiesPrincipalArgsForCall)
}

func (fake *SigningIdentity) SatisfiesPrincipalCalls(stub func(*mspa.MSPPrincipal) error) {
	fake.satisfiesPrincipalMutex.Lock()
	defer fake.satisfiesPrincipalMutex.Unlock()
	fake.SatisfiesPrincipalStub = stub
}

func (fake *SigningIdentity) SatisfiesPrincipalArgsForCall(i int) *mspa.MSPPrincipal {
	fake.satisfiesPrincipalMutex.RLock()
	defer fake.satisfiesPrincipalMutex.RUnlock()
	argsForCall := fake.satisfiesPrincipalArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SigningIdentity) SatisfiesPrincipalReturns(result1 error) {
	fake.satisfiesPrincipalMutex.Lock()
	defer fake.satisfiesPrincipalMutex.Unlock()
	fake.SatisfiesPrincipalStub = nil
	fake.satisfiesPrincipalReturns = struct {
		result1 error
	}{result1}
}

func (fake *SigningIdentity) SatisfiesPrincipalReturnsOnCall(i int, result1 error) {
	fake.satisfiesPrincipalMutex.Lock()
	defer fake.satisfiesPrincipalMutex.Unlock()
	fake.SatisfiesPrincipalStub = nil
	if fake.satisfiesPrincipalReturnsOnCall == nil {
		fake.satisfiesPrincipalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.satisfiesPrincipalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SigningIdentity) Serialize() ([]byte, error) {
	fake.serializeMutex.Lock()
	ret, specificReturn := fake.serializeReturnsOnCall[len(fake.serializeArgsForCall)]
	fake.serializeArgsForCall = append(fake.serializeArgsForCall, struct {
	}{})
	stub := fake.SerializeStub
	fakeReturns := fake.serializeReturns
	fake.recordInvocation("Serialize", []interface{}{})
	fake.serializeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SigningIdentity) SerializeCallCount() int {
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	return len(fake.serializeArgsForCall)
}

func (fake *SigningIdentity) SerializeCalls(stub func() ([]byte, error)) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = stub
}

func (fake *SigningIdentity) SerializeReturns(result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	fake.serializeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SigningIdentity) SerializeReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	if fake.serializeReturnsOnCall == nil {
		fake.serializeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.serializeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SigningIdentity) Sign(arg1 []byte) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.signMutex.Lock()
	ret, specificReturn := fake.signReturnsOnCall[len(fake.signArgsForCall)]
	fake.signArgsForCall = append(fake.signArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.SignStub
	fakeReturns := fake.signReturns
	fake.recordInvocation("Sign", []interface{}{arg1Copy})
	fake.signMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SigningIdentity) SignCallCount() int {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return len(fake.signArgsForCall)
}

func (fake *SigningIdentity) SignCalls(stub func([]byte) ([]byte, error)) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = stub
}

func (fake *SigningIdentity) SignArgsForCall(i int) []byte {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	argsForCall := fake.signArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SigningIdentity) SignReturns(result1 []byte, result2 error) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = nil
	fake.signReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SigningIdentity) SignReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = nil
	if fake.signReturnsOnCall == nil {
		fake.signReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.signReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SigningIdentity) Validate() error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
	}{})
	stub := fake.ValidateStub
	fakeReturns := fake.validateReturns
	fake.recordInvocation("Validate", []interface{}{})
	fake.validateMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SigningIdentity) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *SigningIdentity) ValidateCalls(stub func() error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *SigningIdentity) ValidateReturns(result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *SigningIdentity) ValidateReturnsOnCall(i int, result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SigningIdentity) Verify(arg1 []byte, arg2 []byte) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.verifyMutex.Lock()
	ret, specificReturn := fake.verifyReturnsOnCall[len(fake.verifyArgsForCall)]
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		arg1 []byte
		arg2 []byte
	}{arg1Copy, arg2Copy})
	stub := fake.VerifyStub
	fakeReturns := fake.verifyReturns
	fake.recordInvocation("Verify", []interface{}{arg1Copy, arg2Copy})
	fake.verifyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SigningIdentity) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *SigningIdentity) VerifyCalls(stub func([]byte, []byte) error) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = stub
}

func (fake *SigningIdentity) VerifyArgsForCall(i int) ([]byte, []byte) {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	argsForCall := fake.verifyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SigningIdentity) VerifyReturns(result1 error) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 error
	}{result1}
}

func (fake *SigningIdentity) VerifyReturnsOnCall(i int, result1 error) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	if fake.verifyReturnsOnCall == nil {
		fake.verifyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SigningIdentity) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.anonymousMutex.RLock()
	defer fake.anonymousMutex.RUnlock()
	fake.expiresAtMutex.RLock()
	defer fake.expiresAtMutex.RUnlock()
	fake.getIdentifierMutex.RLock()
	defer fake.getIdentifierMutex.RUnlock()
	fake.getMSPIdentifierMutex.RLock()
	defer fake.getMSPIdentifierMutex.RUnlock()
	fake.getOrganizationalUnitsMutex.RLock()
	defer fake.getOrganizationalUnitsMutex.RUnlock()
	fake.getPublicVersionMutex.RLock()
	defer fake.getPublicVersionMutex.RUnlock()
	fake.satisfiesPrincipalMutex.RLock()
	defer fake.satisfiesPrincipalMutex.RUnlock()
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SigningIdentity) recordInvocation(key string, args []interface{}) {
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
