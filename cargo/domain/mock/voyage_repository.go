// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fengye87/dddsample-core/cargo/domain (interfaces: VoyageRepository)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/fengye87/dddsample-core/cargo/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockVoyageRepository is a mock of VoyageRepository interface.
type MockVoyageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVoyageRepositoryMockRecorder
}

// MockVoyageRepositoryMockRecorder is the mock recorder for MockVoyageRepository.
type MockVoyageRepositoryMockRecorder struct {
	mock *MockVoyageRepository
}

// NewMockVoyageRepository creates a new mock instance.
func NewMockVoyageRepository(ctrl *gomock.Controller) *MockVoyageRepository {
	mock := &MockVoyageRepository{ctrl: ctrl}
	mock.recorder = &MockVoyageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVoyageRepository) EXPECT() *MockVoyageRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockVoyageRepository) Find(arg0 domain.VoyageNumber) (*domain.Voyage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*domain.Voyage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockVoyageRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockVoyageRepository)(nil).Find), arg0)
}

// FindAll mocks base method.
func (m *MockVoyageRepository) FindAll() ([]domain.Voyage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]domain.Voyage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockVoyageRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockVoyageRepository)(nil).FindAll))
}