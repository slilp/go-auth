// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/slilp/go-auth/repository/account (interfaces: AccountRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repository "github.com/slilp/go-auth/repository/account"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountRepository) Create(arg0 repository.AccountEntity) (repository.AccountEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(repository.AccountEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockAccountRepository) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAccountRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccountRepository)(nil).Delete), arg0)
}

// GetByUsername mocks base method.
func (m *MockAccountRepository) GetByUsername(arg0 string) (*repository.AccountEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", arg0)
	ret0, _ := ret[0].(*repository.AccountEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockAccountRepositoryMockRecorder) GetByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockAccountRepository)(nil).GetByUsername), arg0)
}

// Update mocks base method.
func (m *MockAccountRepository) Update(arg0 repository.AccountEntity) (repository.AccountEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(repository.AccountEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAccountRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccountRepository)(nil).Update), arg0)
}