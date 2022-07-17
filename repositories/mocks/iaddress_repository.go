// Code generated by MockGen. DO NOT EDIT.
// Source: veterinary-employee/repositories (interfaces: IAddressRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	models "veterinary-employee/models"

	gomock "github.com/golang/mock/gomock"
)

// MockIAddressRepository is a mock of IAddressRepository interface.
type MockIAddressRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAddressRepositoryMockRecorder
}

// MockIAddressRepositoryMockRecorder is the mock recorder for MockIAddressRepository.
type MockIAddressRepositoryMockRecorder struct {
	mock *MockIAddressRepository
}

// NewMockIAddressRepository creates a new mock instance.
func NewMockIAddressRepository(ctrl *gomock.Controller) *MockIAddressRepository {
	mock := &MockIAddressRepository{ctrl: ctrl}
	mock.recorder = &MockIAddressRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAddressRepository) EXPECT() *MockIAddressRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAddressRepository) Create(arg0 context.Context, arg1 models.Address) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIAddressRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAddressRepository)(nil).Create), arg0, arg1)
}

// Get mocks base method.
func (m *MockIAddressRepository) Get(arg0 context.Context, arg1 interface{}) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIAddressRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIAddressRepository)(nil).Get), arg0, arg1)
}

// Update mocks base method.
func (m *MockIAddressRepository) Update(arg0 context.Context, arg1, arg2 interface{}) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIAddressRepositoryMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIAddressRepository)(nil).Update), arg0, arg1, arg2)
}
