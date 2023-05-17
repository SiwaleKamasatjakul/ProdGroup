// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SiwaleK/ProdGroup/repository (interfaces: PaymentConfigRepository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockPaymentConfigRepository is a mock of PaymentConfigRepository interface.
type MockPaymentConfigRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentConfigRepositoryMockRecorder
}

// MockPaymentConfigRepositoryMockRecorder is the mock recorder for MockPaymentConfigRepository.
type MockPaymentConfigRepositoryMockRecorder struct {
	mock *MockPaymentConfigRepository
}

// NewMockPaymentConfigRepository creates a new mock instance.
func NewMockPaymentConfigRepository(ctrl *gomock.Controller) *MockPaymentConfigRepository {
	mock := &MockPaymentConfigRepository{ctrl: ctrl}
	mock.recorder = &MockPaymentConfigRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentConfigRepository) EXPECT() *MockPaymentConfigRepositoryMockRecorder {
	return m.recorder
}

// GetPaymentConfig mocks base method.
func (m *MockPaymentConfigRepository) GetPaymentConfig(arg0 context.Context) ([]db.GetPaymentConfigRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentConfig", arg0)
	ret0, _ := ret[0].([]db.GetPaymentConfigRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentConfig indicates an expected call of GetPaymentConfig.
func (mr *MockPaymentConfigRepositoryMockRecorder) GetPaymentConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentConfig", reflect.TypeOf((*MockPaymentConfigRepository)(nil).GetPaymentConfig), arg0)
}

// GetPaymentMethod mocks base method.
func (m *MockPaymentConfigRepository) GetPaymentMethod(arg0 context.Context) ([]db.PaymentMethod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentMethod", arg0)
	ret0, _ := ret[0].([]db.PaymentMethod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentMethod indicates an expected call of GetPaymentMethod.
func (mr *MockPaymentConfigRepositoryMockRecorder) GetPaymentMethod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentMethod", reflect.TypeOf((*MockPaymentConfigRepository)(nil).GetPaymentMethod), arg0)
}

// GetPosClientMethod mocks base method.
func (m *MockPaymentConfigRepository) GetPosClientMethod(arg0 context.Context, arg1 *string) ([]db.GetPosClientMethodRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosClientMethod", arg0, arg1)
	ret0, _ := ret[0].([]db.GetPosClientMethodRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosClientMethod indicates an expected call of GetPosClientMethod.
func (mr *MockPaymentConfigRepositoryMockRecorder) GetPosClientMethod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosClientMethod", reflect.TypeOf((*MockPaymentConfigRepository)(nil).GetPosClientMethod), arg0, arg1)
}
