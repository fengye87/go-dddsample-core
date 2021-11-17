// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fengye87/dddsample-core/cargo/domain (interfaces: LocationRepository)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/fengye87/dddsample-core/cargo/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockLocationRepository is a mock of LocationRepository interface.
type MockLocationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLocationRepositoryMockRecorder
}

// MockLocationRepositoryMockRecorder is the mock recorder for MockLocationRepository.
type MockLocationRepositoryMockRecorder struct {
	mock *MockLocationRepository
}

// NewMockLocationRepository creates a new mock instance.
func NewMockLocationRepository(ctrl *gomock.Controller) *MockLocationRepository {
	mock := &MockLocationRepository{ctrl: ctrl}
	mock.recorder = &MockLocationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocationRepository) EXPECT() *MockLocationRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockLocationRepository) Find(arg0 domain.UNLocode) (*domain.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*domain.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockLocationRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockLocationRepository)(nil).Find), arg0)
}

// FindAll mocks base method.
func (m *MockLocationRepository) FindAll() ([]domain.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]domain.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockLocationRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockLocationRepository)(nil).FindAll))
}