// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package fabric is a generated GoMock package.
package fabric

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hyperledger/fabric-gateway/pkg/client"
	grpc "google.golang.org/grpc"
)

// MockCommitInt is a mock of CommitInt interface.
type MockCommitInt struct {
	ctrl     *gomock.Controller
	recorder *MockCommitIntMockRecorder
}

// MockCommitIntMockRecorder is the mock recorder for MockCommitInt.
type MockCommitIntMockRecorder struct {
	mock *MockCommitInt
}

// NewMockCommitInt creates a new mock instance.
func NewMockCommitInt(ctrl *gomock.Controller) *MockCommitInt {
	mock := &MockCommitInt{ctrl: ctrl}
	mock.recorder = &MockCommitIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommitInt) EXPECT() *MockCommitIntMockRecorder {
	return m.recorder
}

// TransactionID mocks base method.
func (m *MockCommitInt) TransactionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TransactionID indicates an expected call of TransactionID.
func (mr *MockCommitIntMockRecorder) TransactionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionID", reflect.TypeOf((*MockCommitInt)(nil).TransactionID))
}

// MockTransactionInt is a mock of TransactionInt interface.
type MockTransactionInt struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionIntMockRecorder
}

// MockTransactionIntMockRecorder is the mock recorder for MockTransactionInt.
type MockTransactionIntMockRecorder struct {
	mock *MockTransactionInt
}

// NewMockTransactionInt creates a new mock instance.
func NewMockTransactionInt(ctrl *gomock.Controller) *MockTransactionInt {
	mock := &MockTransactionInt{ctrl: ctrl}
	mock.recorder = &MockTransactionIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionInt) EXPECT() *MockTransactionIntMockRecorder {
	return m.recorder
}

// Result mocks base method.
func (m *MockTransactionInt) Result() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Result indicates an expected call of Result.
func (mr *MockTransactionIntMockRecorder) Result() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockTransactionInt)(nil).Result))
}

// Submit mocks base method.
func (m *MockTransactionInt) Submit(opts ...grpc.CallOption) (CommitInt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Submit", varargs...)
	ret0, _ := ret[0].(CommitInt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Submit indicates an expected call of Submit.
func (mr *MockTransactionIntMockRecorder) Submit(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Submit", reflect.TypeOf((*MockTransactionInt)(nil).Submit), opts...)
}

// MockProposalInt is a mock of ProposalInt interface.
type MockProposalInt struct {
	ctrl     *gomock.Controller
	recorder *MockProposalIntMockRecorder
}

// MockProposalIntMockRecorder is the mock recorder for MockProposalInt.
type MockProposalIntMockRecorder struct {
	mock *MockProposalInt
}

// NewMockProposalInt creates a new mock instance.
func NewMockProposalInt(ctrl *gomock.Controller) *MockProposalInt {
	mock := &MockProposalInt{ctrl: ctrl}
	mock.recorder = &MockProposalIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProposalInt) EXPECT() *MockProposalIntMockRecorder {
	return m.recorder
}

// Endorse mocks base method.
func (m *MockProposalInt) Endorse(opts ...grpc.CallOption) (TransactionInt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Endorse", varargs...)
	ret0, _ := ret[0].(TransactionInt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Endorse indicates an expected call of Endorse.
func (mr *MockProposalIntMockRecorder) Endorse(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endorse", reflect.TypeOf((*MockProposalInt)(nil).Endorse), opts...)
}

// MockContractInt is a mock of ContractInt interface.
type MockContractInt struct {
	ctrl     *gomock.Controller
	recorder *MockContractIntMockRecorder
}

// MockContractIntMockRecorder is the mock recorder for MockContractInt.
type MockContractIntMockRecorder struct {
	mock *MockContractInt
}

// NewMockContractInt creates a new mock instance.
func NewMockContractInt(ctrl *gomock.Controller) *MockContractInt {
	mock := &MockContractInt{ctrl: ctrl}
	mock.recorder = &MockContractIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractInt) EXPECT() *MockContractIntMockRecorder {
	return m.recorder
}

// EvaluateTransaction mocks base method.
func (m *MockContractInt) EvaluateTransaction(name string, args ...string) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvaluateTransaction", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvaluateTransaction indicates an expected call of EvaluateTransaction.
func (mr *MockContractIntMockRecorder) EvaluateTransaction(name interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvaluateTransaction", reflect.TypeOf((*MockContractInt)(nil).EvaluateTransaction), varargs...)
}

// NewProposal mocks base method.
func (m *MockContractInt) NewProposal(transactionName string, options ...client.ProposalOption) (ProposalInt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{transactionName}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewProposal", varargs...)
	ret0, _ := ret[0].(ProposalInt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewProposal indicates an expected call of NewProposal.
func (mr *MockContractIntMockRecorder) NewProposal(transactionName interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{transactionName}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProposal", reflect.TypeOf((*MockContractInt)(nil).NewProposal), varargs...)
}

// MockNetworkInt is a mock of NetworkInt interface.
type MockNetworkInt struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkIntMockRecorder
}

// MockNetworkIntMockRecorder is the mock recorder for MockNetworkInt.
type MockNetworkIntMockRecorder struct {
	mock *MockNetworkInt
}

// NewMockNetworkInt creates a new mock instance.
func NewMockNetworkInt(ctrl *gomock.Controller) *MockNetworkInt {
	mock := &MockNetworkInt{ctrl: ctrl}
	mock.recorder = &MockNetworkIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkInt) EXPECT() *MockNetworkIntMockRecorder {
	return m.recorder
}

// GetContract mocks base method.
func (m *MockNetworkInt) GetContract(chaincodeName string) ContractInt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract", chaincodeName)
	ret0, _ := ret[0].(ContractInt)
	return ret0
}

// GetContract indicates an expected call of GetContract.
func (mr *MockNetworkIntMockRecorder) GetContract(chaincodeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockNetworkInt)(nil).GetContract), chaincodeName)
}

// MockGatewayInt is a mock of GatewayInt interface.
type MockGatewayInt struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayIntMockRecorder
}

// MockGatewayIntMockRecorder is the mock recorder for MockGatewayInt.
type MockGatewayIntMockRecorder struct {
	mock *MockGatewayInt
}

// NewMockGatewayInt creates a new mock instance.
func NewMockGatewayInt(ctrl *gomock.Controller) *MockGatewayInt {
	mock := &MockGatewayInt{ctrl: ctrl}
	mock.recorder = &MockGatewayIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayInt) EXPECT() *MockGatewayIntMockRecorder {
	return m.recorder
}

// GetNetwork mocks base method.
func (m *MockGatewayInt) GetNetwork(name string) NetworkInt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetwork", name)
	ret0, _ := ret[0].(NetworkInt)
	return ret0
}

// GetNetwork indicates an expected call of GetNetwork.
func (mr *MockGatewayIntMockRecorder) GetNetwork(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetwork", reflect.TypeOf((*MockGatewayInt)(nil).GetNetwork), name)
}
