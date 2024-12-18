// Code generated by MockGen. DO NOT EDIT.
// Source: api/service/deposit_service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	request "simple_bank_solid/model/web/request"
	response "simple_bank_solid/model/web/response"

	gomock "github.com/golang/mock/gomock"
)

// MockDepositServie is a mock of DepositServie interface.
type MockDepositServie struct {
	ctrl     *gomock.Controller
	recorder *MockDepositServieMockRecorder
}

// MockDepositServieMockRecorder is the mock recorder for MockDepositServie.
type MockDepositServieMockRecorder struct {
	mock *MockDepositServie
}

// NewMockDepositServie creates a new mock instance.
func NewMockDepositServie(ctrl *gomock.Controller) *MockDepositServie {
	mock := &MockDepositServie{ctrl: ctrl}
	mock.recorder = &MockDepositServieMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDepositServie) EXPECT() *MockDepositServieMockRecorder {
	return m.recorder
}

// CreateDeposit mocks base method.
func (m *MockDepositServie) CreateDeposit(req request.DepositRequest, userId int64) (response.DepositResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeposit", req, userId)
	ret0, _ := ret[0].(response.DepositResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeposit indicates an expected call of CreateDeposit.
func (mr *MockDepositServieMockRecorder) CreateDeposit(req, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeposit", reflect.TypeOf((*MockDepositServie)(nil).CreateDeposit), req, userId)
}

// Delete mocks base method.
func (m *MockDepositServie) Delete(depositId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", depositId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDepositServieMockRecorder) Delete(depositId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDepositServie)(nil).Delete), depositId)
}

// FetchAllDeposit mocks base method.
func (m *MockDepositServie) FetchAllDeposit() ([]response.DepositResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllDeposit")
	ret0, _ := ret[0].([]response.DepositResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllDeposit indicates an expected call of FetchAllDeposit.
func (mr *MockDepositServieMockRecorder) FetchAllDeposit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllDeposit", reflect.TypeOf((*MockDepositServie)(nil).FetchAllDeposit))
}

// FetchAllDepositByUserId mocks base method.
func (m *MockDepositServie) FetchAllDepositByUserId(userId int64) ([]response.DepositResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllDepositByUserId", userId)
	ret0, _ := ret[0].([]response.DepositResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllDepositByUserId indicates an expected call of FetchAllDepositByUserId.
func (mr *MockDepositServieMockRecorder) FetchAllDepositByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllDepositByUserId", reflect.TypeOf((*MockDepositServie)(nil).FetchAllDepositByUserId), userId)
}

// FetchDepositById mocks base method.
func (m *MockDepositServie) FetchDepositById(DepositId int64) (response.DepositResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDepositById", DepositId)
	ret0, _ := ret[0].(response.DepositResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchDepositById indicates an expected call of FetchDepositById.
func (mr *MockDepositServieMockRecorder) FetchDepositById(DepositId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDepositById", reflect.TypeOf((*MockDepositServie)(nil).FetchDepositById), DepositId)
}
