// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/SmartBFT-Go/fabric/common/chaincode"
	"github.com/SmartBFT-Go/fabric/core/chaincode/lifecycle"
)

type SCCFunctions struct {
	ApproveChaincodeDefinitionForOrgStub        func(string, string, *lifecycle.ChaincodeDefinition, string, lifecycle.ReadableState, lifecycle.ReadWritableState) error
	approveChaincodeDefinitionForOrgMutex       sync.RWMutex
	approveChaincodeDefinitionForOrgArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 string
		arg5 lifecycle.ReadableState
		arg6 lifecycle.ReadWritableState
	}
	approveChaincodeDefinitionForOrgReturns struct {
		result1 error
	}
	approveChaincodeDefinitionForOrgReturnsOnCall map[int]struct {
		result1 error
	}
	CheckCommitReadinessStub        func(string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) (map[string]bool, error)
	checkCommitReadinessMutex       sync.RWMutex
	checkCommitReadinessArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 lifecycle.ReadWritableState
		arg5 []lifecycle.OpaqueState
	}
	checkCommitReadinessReturns struct {
		result1 map[string]bool
		result2 error
	}
	checkCommitReadinessReturnsOnCall map[int]struct {
		result1 map[string]bool
		result2 error
	}
	CommitChaincodeDefinitionStub        func(string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) (map[string]bool, error)
	commitChaincodeDefinitionMutex       sync.RWMutex
	commitChaincodeDefinitionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 lifecycle.ReadWritableState
		arg5 []lifecycle.OpaqueState
	}
	commitChaincodeDefinitionReturns struct {
		result1 map[string]bool
		result2 error
	}
	commitChaincodeDefinitionReturnsOnCall map[int]struct {
		result1 map[string]bool
		result2 error
	}
	GetInstalledChaincodePackageStub        func(string) ([]byte, error)
	getInstalledChaincodePackageMutex       sync.RWMutex
	getInstalledChaincodePackageArgsForCall []struct {
		arg1 string
	}
	getInstalledChaincodePackageReturns struct {
		result1 []byte
		result2 error
	}
	getInstalledChaincodePackageReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	InstallChaincodeStub        func([]byte) (*chaincode.InstalledChaincode, error)
	installChaincodeMutex       sync.RWMutex
	installChaincodeArgsForCall []struct {
		arg1 []byte
	}
	installChaincodeReturns struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}
	installChaincodeReturnsOnCall map[int]struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}
	QueryApprovedChaincodeDefinitionStub        func(string, string, int64, lifecycle.ReadableState, lifecycle.ReadableState) (*lifecycle.ApprovedChaincodeDefinition, error)
	queryApprovedChaincodeDefinitionMutex       sync.RWMutex
	queryApprovedChaincodeDefinitionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int64
		arg4 lifecycle.ReadableState
		arg5 lifecycle.ReadableState
	}
	queryApprovedChaincodeDefinitionReturns struct {
		result1 *lifecycle.ApprovedChaincodeDefinition
		result2 error
	}
	queryApprovedChaincodeDefinitionReturnsOnCall map[int]struct {
		result1 *lifecycle.ApprovedChaincodeDefinition
		result2 error
	}
	QueryChaincodeDefinitionStub        func(string, lifecycle.ReadableState) (*lifecycle.ChaincodeDefinition, error)
	queryChaincodeDefinitionMutex       sync.RWMutex
	queryChaincodeDefinitionArgsForCall []struct {
		arg1 string
		arg2 lifecycle.ReadableState
	}
	queryChaincodeDefinitionReturns struct {
		result1 *lifecycle.ChaincodeDefinition
		result2 error
	}
	queryChaincodeDefinitionReturnsOnCall map[int]struct {
		result1 *lifecycle.ChaincodeDefinition
		result2 error
	}
	QueryInstalledChaincodeStub        func(string) (*chaincode.InstalledChaincode, error)
	queryInstalledChaincodeMutex       sync.RWMutex
	queryInstalledChaincodeArgsForCall []struct {
		arg1 string
	}
	queryInstalledChaincodeReturns struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}
	queryInstalledChaincodeReturnsOnCall map[int]struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}
	QueryInstalledChaincodesStub        func() []*chaincode.InstalledChaincode
	queryInstalledChaincodesMutex       sync.RWMutex
	queryInstalledChaincodesArgsForCall []struct {
	}
	queryInstalledChaincodesReturns struct {
		result1 []*chaincode.InstalledChaincode
	}
	queryInstalledChaincodesReturnsOnCall map[int]struct {
		result1 []*chaincode.InstalledChaincode
	}
	QueryNamespaceDefinitionsStub        func(lifecycle.RangeableState) (map[string]string, error)
	queryNamespaceDefinitionsMutex       sync.RWMutex
	queryNamespaceDefinitionsArgsForCall []struct {
		arg1 lifecycle.RangeableState
	}
	queryNamespaceDefinitionsReturns struct {
		result1 map[string]string
		result2 error
	}
	queryNamespaceDefinitionsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	QueryOrgApprovalsStub        func(string, *lifecycle.ChaincodeDefinition, []lifecycle.OpaqueState) (map[string]bool, error)
	queryOrgApprovalsMutex       sync.RWMutex
	queryOrgApprovalsArgsForCall []struct {
		arg1 string
		arg2 *lifecycle.ChaincodeDefinition
		arg3 []lifecycle.OpaqueState
	}
	queryOrgApprovalsReturns struct {
		result1 map[string]bool
		result2 error
	}
	queryOrgApprovalsReturnsOnCall map[int]struct {
		result1 map[string]bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrg(arg1 string, arg2 string, arg3 *lifecycle.ChaincodeDefinition, arg4 string, arg5 lifecycle.ReadableState, arg6 lifecycle.ReadWritableState) error {
	fake.approveChaincodeDefinitionForOrgMutex.Lock()
	ret, specificReturn := fake.approveChaincodeDefinitionForOrgReturnsOnCall[len(fake.approveChaincodeDefinitionForOrgArgsForCall)]
	fake.approveChaincodeDefinitionForOrgArgsForCall = append(fake.approveChaincodeDefinitionForOrgArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 string
		arg5 lifecycle.ReadableState
		arg6 lifecycle.ReadWritableState
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	stub := fake.ApproveChaincodeDefinitionForOrgStub
	fakeReturns := fake.approveChaincodeDefinitionForOrgReturns
	fake.recordInvocation("ApproveChaincodeDefinitionForOrg", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.approveChaincodeDefinitionForOrgMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgCallCount() int {
	fake.approveChaincodeDefinitionForOrgMutex.RLock()
	defer fake.approveChaincodeDefinitionForOrgMutex.RUnlock()
	return len(fake.approveChaincodeDefinitionForOrgArgsForCall)
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgCalls(stub func(string, string, *lifecycle.ChaincodeDefinition, string, lifecycle.ReadableState, lifecycle.ReadWritableState) error) {
	fake.approveChaincodeDefinitionForOrgMutex.Lock()
	defer fake.approveChaincodeDefinitionForOrgMutex.Unlock()
	fake.ApproveChaincodeDefinitionForOrgStub = stub
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgArgsForCall(i int) (string, string, *lifecycle.ChaincodeDefinition, string, lifecycle.ReadableState, lifecycle.ReadWritableState) {
	fake.approveChaincodeDefinitionForOrgMutex.RLock()
	defer fake.approveChaincodeDefinitionForOrgMutex.RUnlock()
	argsForCall := fake.approveChaincodeDefinitionForOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgReturns(result1 error) {
	fake.approveChaincodeDefinitionForOrgMutex.Lock()
	defer fake.approveChaincodeDefinitionForOrgMutex.Unlock()
	fake.ApproveChaincodeDefinitionForOrgStub = nil
	fake.approveChaincodeDefinitionForOrgReturns = struct {
		result1 error
	}{result1}
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgReturnsOnCall(i int, result1 error) {
	fake.approveChaincodeDefinitionForOrgMutex.Lock()
	defer fake.approveChaincodeDefinitionForOrgMutex.Unlock()
	fake.ApproveChaincodeDefinitionForOrgStub = nil
	if fake.approveChaincodeDefinitionForOrgReturnsOnCall == nil {
		fake.approveChaincodeDefinitionForOrgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.approveChaincodeDefinitionForOrgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SCCFunctions) CheckCommitReadiness(arg1 string, arg2 string, arg3 *lifecycle.ChaincodeDefinition, arg4 lifecycle.ReadWritableState, arg5 []lifecycle.OpaqueState) (map[string]bool, error) {
	var arg5Copy []lifecycle.OpaqueState
	if arg5 != nil {
		arg5Copy = make([]lifecycle.OpaqueState, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.checkCommitReadinessMutex.Lock()
	ret, specificReturn := fake.checkCommitReadinessReturnsOnCall[len(fake.checkCommitReadinessArgsForCall)]
	fake.checkCommitReadinessArgsForCall = append(fake.checkCommitReadinessArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 lifecycle.ReadWritableState
		arg5 []lifecycle.OpaqueState
	}{arg1, arg2, arg3, arg4, arg5Copy})
	stub := fake.CheckCommitReadinessStub
	fakeReturns := fake.checkCommitReadinessReturns
	fake.recordInvocation("CheckCommitReadiness", []interface{}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.checkCommitReadinessMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) CheckCommitReadinessCallCount() int {
	fake.checkCommitReadinessMutex.RLock()
	defer fake.checkCommitReadinessMutex.RUnlock()
	return len(fake.checkCommitReadinessArgsForCall)
}

func (fake *SCCFunctions) CheckCommitReadinessCalls(stub func(string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) (map[string]bool, error)) {
	fake.checkCommitReadinessMutex.Lock()
	defer fake.checkCommitReadinessMutex.Unlock()
	fake.CheckCommitReadinessStub = stub
}

func (fake *SCCFunctions) CheckCommitReadinessArgsForCall(i int) (string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) {
	fake.checkCommitReadinessMutex.RLock()
	defer fake.checkCommitReadinessMutex.RUnlock()
	argsForCall := fake.checkCommitReadinessArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *SCCFunctions) CheckCommitReadinessReturns(result1 map[string]bool, result2 error) {
	fake.checkCommitReadinessMutex.Lock()
	defer fake.checkCommitReadinessMutex.Unlock()
	fake.CheckCommitReadinessStub = nil
	fake.checkCommitReadinessReturns = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) CheckCommitReadinessReturnsOnCall(i int, result1 map[string]bool, result2 error) {
	fake.checkCommitReadinessMutex.Lock()
	defer fake.checkCommitReadinessMutex.Unlock()
	fake.CheckCommitReadinessStub = nil
	if fake.checkCommitReadinessReturnsOnCall == nil {
		fake.checkCommitReadinessReturnsOnCall = make(map[int]struct {
			result1 map[string]bool
			result2 error
		})
	}
	fake.checkCommitReadinessReturnsOnCall[i] = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) CommitChaincodeDefinition(arg1 string, arg2 string, arg3 *lifecycle.ChaincodeDefinition, arg4 lifecycle.ReadWritableState, arg5 []lifecycle.OpaqueState) (map[string]bool, error) {
	var arg5Copy []lifecycle.OpaqueState
	if arg5 != nil {
		arg5Copy = make([]lifecycle.OpaqueState, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.commitChaincodeDefinitionMutex.Lock()
	ret, specificReturn := fake.commitChaincodeDefinitionReturnsOnCall[len(fake.commitChaincodeDefinitionArgsForCall)]
	fake.commitChaincodeDefinitionArgsForCall = append(fake.commitChaincodeDefinitionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 lifecycle.ReadWritableState
		arg5 []lifecycle.OpaqueState
	}{arg1, arg2, arg3, arg4, arg5Copy})
	stub := fake.CommitChaincodeDefinitionStub
	fakeReturns := fake.commitChaincodeDefinitionReturns
	fake.recordInvocation("CommitChaincodeDefinition", []interface{}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.commitChaincodeDefinitionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) CommitChaincodeDefinitionCallCount() int {
	fake.commitChaincodeDefinitionMutex.RLock()
	defer fake.commitChaincodeDefinitionMutex.RUnlock()
	return len(fake.commitChaincodeDefinitionArgsForCall)
}

func (fake *SCCFunctions) CommitChaincodeDefinitionCalls(stub func(string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) (map[string]bool, error)) {
	fake.commitChaincodeDefinitionMutex.Lock()
	defer fake.commitChaincodeDefinitionMutex.Unlock()
	fake.CommitChaincodeDefinitionStub = stub
}

func (fake *SCCFunctions) CommitChaincodeDefinitionArgsForCall(i int) (string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) {
	fake.commitChaincodeDefinitionMutex.RLock()
	defer fake.commitChaincodeDefinitionMutex.RUnlock()
	argsForCall := fake.commitChaincodeDefinitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *SCCFunctions) CommitChaincodeDefinitionReturns(result1 map[string]bool, result2 error) {
	fake.commitChaincodeDefinitionMutex.Lock()
	defer fake.commitChaincodeDefinitionMutex.Unlock()
	fake.CommitChaincodeDefinitionStub = nil
	fake.commitChaincodeDefinitionReturns = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) CommitChaincodeDefinitionReturnsOnCall(i int, result1 map[string]bool, result2 error) {
	fake.commitChaincodeDefinitionMutex.Lock()
	defer fake.commitChaincodeDefinitionMutex.Unlock()
	fake.CommitChaincodeDefinitionStub = nil
	if fake.commitChaincodeDefinitionReturnsOnCall == nil {
		fake.commitChaincodeDefinitionReturnsOnCall = make(map[int]struct {
			result1 map[string]bool
			result2 error
		})
	}
	fake.commitChaincodeDefinitionReturnsOnCall[i] = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) GetInstalledChaincodePackage(arg1 string) ([]byte, error) {
	fake.getInstalledChaincodePackageMutex.Lock()
	ret, specificReturn := fake.getInstalledChaincodePackageReturnsOnCall[len(fake.getInstalledChaincodePackageArgsForCall)]
	fake.getInstalledChaincodePackageArgsForCall = append(fake.getInstalledChaincodePackageArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetInstalledChaincodePackageStub
	fakeReturns := fake.getInstalledChaincodePackageReturns
	fake.recordInvocation("GetInstalledChaincodePackage", []interface{}{arg1})
	fake.getInstalledChaincodePackageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) GetInstalledChaincodePackageCallCount() int {
	fake.getInstalledChaincodePackageMutex.RLock()
	defer fake.getInstalledChaincodePackageMutex.RUnlock()
	return len(fake.getInstalledChaincodePackageArgsForCall)
}

func (fake *SCCFunctions) GetInstalledChaincodePackageCalls(stub func(string) ([]byte, error)) {
	fake.getInstalledChaincodePackageMutex.Lock()
	defer fake.getInstalledChaincodePackageMutex.Unlock()
	fake.GetInstalledChaincodePackageStub = stub
}

func (fake *SCCFunctions) GetInstalledChaincodePackageArgsForCall(i int) string {
	fake.getInstalledChaincodePackageMutex.RLock()
	defer fake.getInstalledChaincodePackageMutex.RUnlock()
	argsForCall := fake.getInstalledChaincodePackageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SCCFunctions) GetInstalledChaincodePackageReturns(result1 []byte, result2 error) {
	fake.getInstalledChaincodePackageMutex.Lock()
	defer fake.getInstalledChaincodePackageMutex.Unlock()
	fake.GetInstalledChaincodePackageStub = nil
	fake.getInstalledChaincodePackageReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) GetInstalledChaincodePackageReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getInstalledChaincodePackageMutex.Lock()
	defer fake.getInstalledChaincodePackageMutex.Unlock()
	fake.GetInstalledChaincodePackageStub = nil
	if fake.getInstalledChaincodePackageReturnsOnCall == nil {
		fake.getInstalledChaincodePackageReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getInstalledChaincodePackageReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) InstallChaincode(arg1 []byte) (*chaincode.InstalledChaincode, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.installChaincodeMutex.Lock()
	ret, specificReturn := fake.installChaincodeReturnsOnCall[len(fake.installChaincodeArgsForCall)]
	fake.installChaincodeArgsForCall = append(fake.installChaincodeArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.InstallChaincodeStub
	fakeReturns := fake.installChaincodeReturns
	fake.recordInvocation("InstallChaincode", []interface{}{arg1Copy})
	fake.installChaincodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) InstallChaincodeCallCount() int {
	fake.installChaincodeMutex.RLock()
	defer fake.installChaincodeMutex.RUnlock()
	return len(fake.installChaincodeArgsForCall)
}

func (fake *SCCFunctions) InstallChaincodeCalls(stub func([]byte) (*chaincode.InstalledChaincode, error)) {
	fake.installChaincodeMutex.Lock()
	defer fake.installChaincodeMutex.Unlock()
	fake.InstallChaincodeStub = stub
}

func (fake *SCCFunctions) InstallChaincodeArgsForCall(i int) []byte {
	fake.installChaincodeMutex.RLock()
	defer fake.installChaincodeMutex.RUnlock()
	argsForCall := fake.installChaincodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SCCFunctions) InstallChaincodeReturns(result1 *chaincode.InstalledChaincode, result2 error) {
	fake.installChaincodeMutex.Lock()
	defer fake.installChaincodeMutex.Unlock()
	fake.InstallChaincodeStub = nil
	fake.installChaincodeReturns = struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) InstallChaincodeReturnsOnCall(i int, result1 *chaincode.InstalledChaincode, result2 error) {
	fake.installChaincodeMutex.Lock()
	defer fake.installChaincodeMutex.Unlock()
	fake.InstallChaincodeStub = nil
	if fake.installChaincodeReturnsOnCall == nil {
		fake.installChaincodeReturnsOnCall = make(map[int]struct {
			result1 *chaincode.InstalledChaincode
			result2 error
		})
	}
	fake.installChaincodeReturnsOnCall[i] = struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryApprovedChaincodeDefinition(arg1 string, arg2 string, arg3 int64, arg4 lifecycle.ReadableState, arg5 lifecycle.ReadableState) (*lifecycle.ApprovedChaincodeDefinition, error) {
	fake.queryApprovedChaincodeDefinitionMutex.Lock()
	ret, specificReturn := fake.queryApprovedChaincodeDefinitionReturnsOnCall[len(fake.queryApprovedChaincodeDefinitionArgsForCall)]
	fake.queryApprovedChaincodeDefinitionArgsForCall = append(fake.queryApprovedChaincodeDefinitionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int64
		arg4 lifecycle.ReadableState
		arg5 lifecycle.ReadableState
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.QueryApprovedChaincodeDefinitionStub
	fakeReturns := fake.queryApprovedChaincodeDefinitionReturns
	fake.recordInvocation("QueryApprovedChaincodeDefinition", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.queryApprovedChaincodeDefinitionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryApprovedChaincodeDefinitionCallCount() int {
	fake.queryApprovedChaincodeDefinitionMutex.RLock()
	defer fake.queryApprovedChaincodeDefinitionMutex.RUnlock()
	return len(fake.queryApprovedChaincodeDefinitionArgsForCall)
}

func (fake *SCCFunctions) QueryApprovedChaincodeDefinitionCalls(stub func(string, string, int64, lifecycle.ReadableState, lifecycle.ReadableState) (*lifecycle.ApprovedChaincodeDefinition, error)) {
	fake.queryApprovedChaincodeDefinitionMutex.Lock()
	defer fake.queryApprovedChaincodeDefinitionMutex.Unlock()
	fake.QueryApprovedChaincodeDefinitionStub = stub
}

func (fake *SCCFunctions) QueryApprovedChaincodeDefinitionArgsForCall(i int) (string, string, int64, lifecycle.ReadableState, lifecycle.ReadableState) {
	fake.queryApprovedChaincodeDefinitionMutex.RLock()
	defer fake.queryApprovedChaincodeDefinitionMutex.RUnlock()
	argsForCall := fake.queryApprovedChaincodeDefinitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *SCCFunctions) QueryApprovedChaincodeDefinitionReturns(result1 *lifecycle.ApprovedChaincodeDefinition, result2 error) {
	fake.queryApprovedChaincodeDefinitionMutex.Lock()
	defer fake.queryApprovedChaincodeDefinitionMutex.Unlock()
	fake.QueryApprovedChaincodeDefinitionStub = nil
	fake.queryApprovedChaincodeDefinitionReturns = struct {
		result1 *lifecycle.ApprovedChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryApprovedChaincodeDefinitionReturnsOnCall(i int, result1 *lifecycle.ApprovedChaincodeDefinition, result2 error) {
	fake.queryApprovedChaincodeDefinitionMutex.Lock()
	defer fake.queryApprovedChaincodeDefinitionMutex.Unlock()
	fake.QueryApprovedChaincodeDefinitionStub = nil
	if fake.queryApprovedChaincodeDefinitionReturnsOnCall == nil {
		fake.queryApprovedChaincodeDefinitionReturnsOnCall = make(map[int]struct {
			result1 *lifecycle.ApprovedChaincodeDefinition
			result2 error
		})
	}
	fake.queryApprovedChaincodeDefinitionReturnsOnCall[i] = struct {
		result1 *lifecycle.ApprovedChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryChaincodeDefinition(arg1 string, arg2 lifecycle.ReadableState) (*lifecycle.ChaincodeDefinition, error) {
	fake.queryChaincodeDefinitionMutex.Lock()
	ret, specificReturn := fake.queryChaincodeDefinitionReturnsOnCall[len(fake.queryChaincodeDefinitionArgsForCall)]
	fake.queryChaincodeDefinitionArgsForCall = append(fake.queryChaincodeDefinitionArgsForCall, struct {
		arg1 string
		arg2 lifecycle.ReadableState
	}{arg1, arg2})
	stub := fake.QueryChaincodeDefinitionStub
	fakeReturns := fake.queryChaincodeDefinitionReturns
	fake.recordInvocation("QueryChaincodeDefinition", []interface{}{arg1, arg2})
	fake.queryChaincodeDefinitionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryChaincodeDefinitionCallCount() int {
	fake.queryChaincodeDefinitionMutex.RLock()
	defer fake.queryChaincodeDefinitionMutex.RUnlock()
	return len(fake.queryChaincodeDefinitionArgsForCall)
}

func (fake *SCCFunctions) QueryChaincodeDefinitionCalls(stub func(string, lifecycle.ReadableState) (*lifecycle.ChaincodeDefinition, error)) {
	fake.queryChaincodeDefinitionMutex.Lock()
	defer fake.queryChaincodeDefinitionMutex.Unlock()
	fake.QueryChaincodeDefinitionStub = stub
}

func (fake *SCCFunctions) QueryChaincodeDefinitionArgsForCall(i int) (string, lifecycle.ReadableState) {
	fake.queryChaincodeDefinitionMutex.RLock()
	defer fake.queryChaincodeDefinitionMutex.RUnlock()
	argsForCall := fake.queryChaincodeDefinitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SCCFunctions) QueryChaincodeDefinitionReturns(result1 *lifecycle.ChaincodeDefinition, result2 error) {
	fake.queryChaincodeDefinitionMutex.Lock()
	defer fake.queryChaincodeDefinitionMutex.Unlock()
	fake.QueryChaincodeDefinitionStub = nil
	fake.queryChaincodeDefinitionReturns = struct {
		result1 *lifecycle.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryChaincodeDefinitionReturnsOnCall(i int, result1 *lifecycle.ChaincodeDefinition, result2 error) {
	fake.queryChaincodeDefinitionMutex.Lock()
	defer fake.queryChaincodeDefinitionMutex.Unlock()
	fake.QueryChaincodeDefinitionStub = nil
	if fake.queryChaincodeDefinitionReturnsOnCall == nil {
		fake.queryChaincodeDefinitionReturnsOnCall = make(map[int]struct {
			result1 *lifecycle.ChaincodeDefinition
			result2 error
		})
	}
	fake.queryChaincodeDefinitionReturnsOnCall[i] = struct {
		result1 *lifecycle.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryInstalledChaincode(arg1 string) (*chaincode.InstalledChaincode, error) {
	fake.queryInstalledChaincodeMutex.Lock()
	ret, specificReturn := fake.queryInstalledChaincodeReturnsOnCall[len(fake.queryInstalledChaincodeArgsForCall)]
	fake.queryInstalledChaincodeArgsForCall = append(fake.queryInstalledChaincodeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.QueryInstalledChaincodeStub
	fakeReturns := fake.queryInstalledChaincodeReturns
	fake.recordInvocation("QueryInstalledChaincode", []interface{}{arg1})
	fake.queryInstalledChaincodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryInstalledChaincodeCallCount() int {
	fake.queryInstalledChaincodeMutex.RLock()
	defer fake.queryInstalledChaincodeMutex.RUnlock()
	return len(fake.queryInstalledChaincodeArgsForCall)
}

func (fake *SCCFunctions) QueryInstalledChaincodeCalls(stub func(string) (*chaincode.InstalledChaincode, error)) {
	fake.queryInstalledChaincodeMutex.Lock()
	defer fake.queryInstalledChaincodeMutex.Unlock()
	fake.QueryInstalledChaincodeStub = stub
}

func (fake *SCCFunctions) QueryInstalledChaincodeArgsForCall(i int) string {
	fake.queryInstalledChaincodeMutex.RLock()
	defer fake.queryInstalledChaincodeMutex.RUnlock()
	argsForCall := fake.queryInstalledChaincodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SCCFunctions) QueryInstalledChaincodeReturns(result1 *chaincode.InstalledChaincode, result2 error) {
	fake.queryInstalledChaincodeMutex.Lock()
	defer fake.queryInstalledChaincodeMutex.Unlock()
	fake.QueryInstalledChaincodeStub = nil
	fake.queryInstalledChaincodeReturns = struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryInstalledChaincodeReturnsOnCall(i int, result1 *chaincode.InstalledChaincode, result2 error) {
	fake.queryInstalledChaincodeMutex.Lock()
	defer fake.queryInstalledChaincodeMutex.Unlock()
	fake.QueryInstalledChaincodeStub = nil
	if fake.queryInstalledChaincodeReturnsOnCall == nil {
		fake.queryInstalledChaincodeReturnsOnCall = make(map[int]struct {
			result1 *chaincode.InstalledChaincode
			result2 error
		})
	}
	fake.queryInstalledChaincodeReturnsOnCall[i] = struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryInstalledChaincodes() []*chaincode.InstalledChaincode {
	fake.queryInstalledChaincodesMutex.Lock()
	ret, specificReturn := fake.queryInstalledChaincodesReturnsOnCall[len(fake.queryInstalledChaincodesArgsForCall)]
	fake.queryInstalledChaincodesArgsForCall = append(fake.queryInstalledChaincodesArgsForCall, struct {
	}{})
	stub := fake.QueryInstalledChaincodesStub
	fakeReturns := fake.queryInstalledChaincodesReturns
	fake.recordInvocation("QueryInstalledChaincodes", []interface{}{})
	fake.queryInstalledChaincodesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *SCCFunctions) QueryInstalledChaincodesCallCount() int {
	fake.queryInstalledChaincodesMutex.RLock()
	defer fake.queryInstalledChaincodesMutex.RUnlock()
	return len(fake.queryInstalledChaincodesArgsForCall)
}

func (fake *SCCFunctions) QueryInstalledChaincodesCalls(stub func() []*chaincode.InstalledChaincode) {
	fake.queryInstalledChaincodesMutex.Lock()
	defer fake.queryInstalledChaincodesMutex.Unlock()
	fake.QueryInstalledChaincodesStub = stub
}

func (fake *SCCFunctions) QueryInstalledChaincodesReturns(result1 []*chaincode.InstalledChaincode) {
	fake.queryInstalledChaincodesMutex.Lock()
	defer fake.queryInstalledChaincodesMutex.Unlock()
	fake.QueryInstalledChaincodesStub = nil
	fake.queryInstalledChaincodesReturns = struct {
		result1 []*chaincode.InstalledChaincode
	}{result1}
}

func (fake *SCCFunctions) QueryInstalledChaincodesReturnsOnCall(i int, result1 []*chaincode.InstalledChaincode) {
	fake.queryInstalledChaincodesMutex.Lock()
	defer fake.queryInstalledChaincodesMutex.Unlock()
	fake.QueryInstalledChaincodesStub = nil
	if fake.queryInstalledChaincodesReturnsOnCall == nil {
		fake.queryInstalledChaincodesReturnsOnCall = make(map[int]struct {
			result1 []*chaincode.InstalledChaincode
		})
	}
	fake.queryInstalledChaincodesReturnsOnCall[i] = struct {
		result1 []*chaincode.InstalledChaincode
	}{result1}
}

func (fake *SCCFunctions) QueryNamespaceDefinitions(arg1 lifecycle.RangeableState) (map[string]string, error) {
	fake.queryNamespaceDefinitionsMutex.Lock()
	ret, specificReturn := fake.queryNamespaceDefinitionsReturnsOnCall[len(fake.queryNamespaceDefinitionsArgsForCall)]
	fake.queryNamespaceDefinitionsArgsForCall = append(fake.queryNamespaceDefinitionsArgsForCall, struct {
		arg1 lifecycle.RangeableState
	}{arg1})
	stub := fake.QueryNamespaceDefinitionsStub
	fakeReturns := fake.queryNamespaceDefinitionsReturns
	fake.recordInvocation("QueryNamespaceDefinitions", []interface{}{arg1})
	fake.queryNamespaceDefinitionsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsCallCount() int {
	fake.queryNamespaceDefinitionsMutex.RLock()
	defer fake.queryNamespaceDefinitionsMutex.RUnlock()
	return len(fake.queryNamespaceDefinitionsArgsForCall)
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsCalls(stub func(lifecycle.RangeableState) (map[string]string, error)) {
	fake.queryNamespaceDefinitionsMutex.Lock()
	defer fake.queryNamespaceDefinitionsMutex.Unlock()
	fake.QueryNamespaceDefinitionsStub = stub
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsArgsForCall(i int) lifecycle.RangeableState {
	fake.queryNamespaceDefinitionsMutex.RLock()
	defer fake.queryNamespaceDefinitionsMutex.RUnlock()
	argsForCall := fake.queryNamespaceDefinitionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsReturns(result1 map[string]string, result2 error) {
	fake.queryNamespaceDefinitionsMutex.Lock()
	defer fake.queryNamespaceDefinitionsMutex.Unlock()
	fake.QueryNamespaceDefinitionsStub = nil
	fake.queryNamespaceDefinitionsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.queryNamespaceDefinitionsMutex.Lock()
	defer fake.queryNamespaceDefinitionsMutex.Unlock()
	fake.QueryNamespaceDefinitionsStub = nil
	if fake.queryNamespaceDefinitionsReturnsOnCall == nil {
		fake.queryNamespaceDefinitionsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.queryNamespaceDefinitionsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryOrgApprovals(arg1 string, arg2 *lifecycle.ChaincodeDefinition, arg3 []lifecycle.OpaqueState) (map[string]bool, error) {
	var arg3Copy []lifecycle.OpaqueState
	if arg3 != nil {
		arg3Copy = make([]lifecycle.OpaqueState, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.queryOrgApprovalsMutex.Lock()
	ret, specificReturn := fake.queryOrgApprovalsReturnsOnCall[len(fake.queryOrgApprovalsArgsForCall)]
	fake.queryOrgApprovalsArgsForCall = append(fake.queryOrgApprovalsArgsForCall, struct {
		arg1 string
		arg2 *lifecycle.ChaincodeDefinition
		arg3 []lifecycle.OpaqueState
	}{arg1, arg2, arg3Copy})
	stub := fake.QueryOrgApprovalsStub
	fakeReturns := fake.queryOrgApprovalsReturns
	fake.recordInvocation("QueryOrgApprovals", []interface{}{arg1, arg2, arg3Copy})
	fake.queryOrgApprovalsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryOrgApprovalsCallCount() int {
	fake.queryOrgApprovalsMutex.RLock()
	defer fake.queryOrgApprovalsMutex.RUnlock()
	return len(fake.queryOrgApprovalsArgsForCall)
}

func (fake *SCCFunctions) QueryOrgApprovalsCalls(stub func(string, *lifecycle.ChaincodeDefinition, []lifecycle.OpaqueState) (map[string]bool, error)) {
	fake.queryOrgApprovalsMutex.Lock()
	defer fake.queryOrgApprovalsMutex.Unlock()
	fake.QueryOrgApprovalsStub = stub
}

func (fake *SCCFunctions) QueryOrgApprovalsArgsForCall(i int) (string, *lifecycle.ChaincodeDefinition, []lifecycle.OpaqueState) {
	fake.queryOrgApprovalsMutex.RLock()
	defer fake.queryOrgApprovalsMutex.RUnlock()
	argsForCall := fake.queryOrgApprovalsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SCCFunctions) QueryOrgApprovalsReturns(result1 map[string]bool, result2 error) {
	fake.queryOrgApprovalsMutex.Lock()
	defer fake.queryOrgApprovalsMutex.Unlock()
	fake.QueryOrgApprovalsStub = nil
	fake.queryOrgApprovalsReturns = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryOrgApprovalsReturnsOnCall(i int, result1 map[string]bool, result2 error) {
	fake.queryOrgApprovalsMutex.Lock()
	defer fake.queryOrgApprovalsMutex.Unlock()
	fake.QueryOrgApprovalsStub = nil
	if fake.queryOrgApprovalsReturnsOnCall == nil {
		fake.queryOrgApprovalsReturnsOnCall = make(map[int]struct {
			result1 map[string]bool
			result2 error
		})
	}
	fake.queryOrgApprovalsReturnsOnCall[i] = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.approveChaincodeDefinitionForOrgMutex.RLock()
	defer fake.approveChaincodeDefinitionForOrgMutex.RUnlock()
	fake.checkCommitReadinessMutex.RLock()
	defer fake.checkCommitReadinessMutex.RUnlock()
	fake.commitChaincodeDefinitionMutex.RLock()
	defer fake.commitChaincodeDefinitionMutex.RUnlock()
	fake.getInstalledChaincodePackageMutex.RLock()
	defer fake.getInstalledChaincodePackageMutex.RUnlock()
	fake.installChaincodeMutex.RLock()
	defer fake.installChaincodeMutex.RUnlock()
	fake.queryApprovedChaincodeDefinitionMutex.RLock()
	defer fake.queryApprovedChaincodeDefinitionMutex.RUnlock()
	fake.queryChaincodeDefinitionMutex.RLock()
	defer fake.queryChaincodeDefinitionMutex.RUnlock()
	fake.queryInstalledChaincodeMutex.RLock()
	defer fake.queryInstalledChaincodeMutex.RUnlock()
	fake.queryInstalledChaincodesMutex.RLock()
	defer fake.queryInstalledChaincodesMutex.RUnlock()
	fake.queryNamespaceDefinitionsMutex.RLock()
	defer fake.queryNamespaceDefinitionsMutex.RUnlock()
	fake.queryOrgApprovalsMutex.RLock()
	defer fake.queryOrgApprovalsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SCCFunctions) recordInvocation(key string, args []interface{}) {
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
