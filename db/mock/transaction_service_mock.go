// Code generated by MockGen. DO NOT EDIT.
// Source: api/service/transaction_service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	request "simple_bank_solid/model/web/request"
	response "simple_bank_solid/model/web/response"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// DeleteTransfer mocks base method.
func (m *MockTransactionService) DeleteTransfer(TxId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransfer", TxId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransfer indicates an expected call of DeleteTransfer.
func (mr *MockTransactionServiceMockRecorder) DeleteTransfer(TxId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransfer", reflect.TypeOf((*MockTransactionService)(nil).DeleteTransfer), TxId)
}

// FecthAllTransfer mocks base method.
func (m *MockTransactionService) FecthAllTransfer() ([]response.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FecthAllTransfer")
	ret0, _ := ret[0].([]response.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FecthAllTransfer indicates an expected call of FecthAllTransfer.
func (mr *MockTransactionServiceMockRecorder) FecthAllTransfer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FecthAllTransfer", reflect.TypeOf((*MockTransactionService)(nil).FecthAllTransfer))
}

// FecthAllTransferByUserId mocks base method.
func (m *MockTransactionService) FecthAllTransferByUserId(UserId int64) ([]response.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FecthAllTransferByUserId", UserId)
	ret0, _ := ret[0].([]response.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FecthAllTransferByUserId indicates an expected call of FecthAllTransferByUserId.
func (mr *MockTransactionServiceMockRecorder) FecthAllTransferByUserId(UserId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FecthAllTransferByUserId", reflect.TypeOf((*MockTransactionService)(nil).FecthAllTransferByUserId), UserId)
}

// FecthTransferById mocks base method.
func (m *MockTransactionService) FecthTransferById(TransactionId string) (response.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FecthTransferById", TransactionId)
	ret0, _ := ret[0].(response.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FecthTransferById indicates an expected call of FecthTransferById.
func (mr *MockTransactionServiceMockRecorder) FecthTransferById(TransactionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FecthTransferById", reflect.TypeOf((*MockTransactionService)(nil).FecthTransferById), TransactionId)
}

// Transfer mocks base method.
func (m *MockTransactionService) Transfer(req request.TransferRequest, userId int64) (response.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", req, userId)
	ret0, _ := ret[0].(response.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockTransactionServiceMockRecorder) Transfer(req, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockTransactionService)(nil).Transfer), req, userId)
}
