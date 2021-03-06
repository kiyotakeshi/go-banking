// Code generated by MockGen. DO NOT EDIT.
// Source: account.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	domain "banking/domain"
	errs "banking/errs"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// FindById mocks base method.
func (m *MockAccountRepository) FindById(accountId string) (*domain.Account, *errs.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", accountId)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(*errs.ApplicationError)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockAccountRepositoryMockRecorder) FindById(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAccountRepository)(nil).FindById), accountId)
}

// Save mocks base method.
func (m *MockAccountRepository) Save(arg0 domain.Account) (*domain.Account, *errs.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(*errs.ApplicationError)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockAccountRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAccountRepository)(nil).Save), arg0)
}

// SaveTransaction mocks base method.
func (m *MockAccountRepository) SaveTransaction(transaction domain.Transaction) (*domain.Transaction, *errs.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTransaction", transaction)
	ret0, _ := ret[0].(*domain.Transaction)
	ret1, _ := ret[1].(*errs.ApplicationError)
	return ret0, ret1
}

// SaveTransaction indicates an expected call of SaveTransaction.
func (mr *MockAccountRepositoryMockRecorder) SaveTransaction(transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTransaction", reflect.TypeOf((*MockAccountRepository)(nil).SaveTransaction), transaction)
}
