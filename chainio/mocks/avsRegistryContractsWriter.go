// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry (interfaces: AvsRegistryWriter)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avsRegistryContractsWriter.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry AvsRegistryWriter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	ecdsa "crypto/ecdsa"
	big "math/big"
	reflect "reflect"

	contractRegistryCoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	bls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsRegistryWriter is a mock of AvsRegistryWriter interface.
type MockAvsRegistryWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAvsRegistryWriterMockRecorder
}

// MockAvsRegistryWriterMockRecorder is the mock recorder for MockAvsRegistryWriter.
type MockAvsRegistryWriterMockRecorder struct {
	mock *MockAvsRegistryWriter
}

// NewMockAvsRegistryWriter creates a new mock instance.
func NewMockAvsRegistryWriter(ctrl *gomock.Controller) *MockAvsRegistryWriter {
	mock := &MockAvsRegistryWriter{ctrl: ctrl}
	mock.recorder = &MockAvsRegistryWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsRegistryWriter) EXPECT() *MockAvsRegistryWriterMockRecorder {
	return m.recorder
}

// DeregisterOperator mocks base method.
func (m *MockAvsRegistryWriter) DeregisterOperator(arg0 context.Context, arg1 []byte, arg2 contractRegistryCoordinator.BN254G1Point) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterOperator", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterOperator indicates an expected call of DeregisterOperator.
func (mr *MockAvsRegistryWriterMockRecorder) DeregisterOperator(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterOperator", reflect.TypeOf((*MockAvsRegistryWriter)(nil).DeregisterOperator), arg0, arg1, arg2)
}

// RegisterOperatorInQuorumWithAVSRegistryCoordinator mocks base method.
func (m *MockAvsRegistryWriter) RegisterOperatorInQuorumWithAVSRegistryCoordinator(arg0 context.Context, arg1 *ecdsa.PrivateKey, arg2 [32]byte, arg3 *big.Int, arg4 *bls.KeyPair, arg5 []byte, arg6 string) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterOperatorInQuorumWithAVSRegistryCoordinator", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterOperatorInQuorumWithAVSRegistryCoordinator indicates an expected call of RegisterOperatorInQuorumWithAVSRegistryCoordinator.
func (mr *MockAvsRegistryWriterMockRecorder) RegisterOperatorInQuorumWithAVSRegistryCoordinator(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOperatorInQuorumWithAVSRegistryCoordinator", reflect.TypeOf((*MockAvsRegistryWriter)(nil).RegisterOperatorInQuorumWithAVSRegistryCoordinator), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// UpdateStakesOfEntireOperatorSetForQuorums mocks base method.
func (m *MockAvsRegistryWriter) UpdateStakesOfEntireOperatorSetForQuorums(arg0 context.Context, arg1 [][]common.Address, arg2 []byte) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStakesOfEntireOperatorSetForQuorums", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStakesOfEntireOperatorSetForQuorums indicates an expected call of UpdateStakesOfEntireOperatorSetForQuorums.
func (mr *MockAvsRegistryWriterMockRecorder) UpdateStakesOfEntireOperatorSetForQuorums(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStakesOfEntireOperatorSetForQuorums", reflect.TypeOf((*MockAvsRegistryWriter)(nil).UpdateStakesOfEntireOperatorSetForQuorums), arg0, arg1, arg2)
}

// UpdateStakesOfOperatorSubsetForAllQuorums mocks base method.
func (m *MockAvsRegistryWriter) UpdateStakesOfOperatorSubsetForAllQuorums(arg0 context.Context, arg1 []common.Address) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStakesOfOperatorSubsetForAllQuorums", arg0, arg1)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStakesOfOperatorSubsetForAllQuorums indicates an expected call of UpdateStakesOfOperatorSubsetForAllQuorums.
func (mr *MockAvsRegistryWriterMockRecorder) UpdateStakesOfOperatorSubsetForAllQuorums(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStakesOfOperatorSubsetForAllQuorums", reflect.TypeOf((*MockAvsRegistryWriter)(nil).UpdateStakesOfOperatorSubsetForAllQuorums), arg0, arg1)
}
