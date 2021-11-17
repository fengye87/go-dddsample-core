// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fengye87/dddsample-core/cargo/application (interfaces: CargoInspectionService)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/fengye87/dddsample-core/cargo/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockCargoInspectionService is a mock of CargoInspectionService interface.
type MockCargoInspectionService struct {
	ctrl     *gomock.Controller
	recorder *MockCargoInspectionServiceMockRecorder
}

// MockCargoInspectionServiceMockRecorder is the mock recorder for MockCargoInspectionService.
type MockCargoInspectionServiceMockRecorder struct {
	mock *MockCargoInspectionService
}

// NewMockCargoInspectionService creates a new mock instance.
func NewMockCargoInspectionService(ctrl *gomock.Controller) *MockCargoInspectionService {
	mock := &MockCargoInspectionService{ctrl: ctrl}
	mock.recorder = &MockCargoInspectionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCargoInspectionService) EXPECT() *MockCargoInspectionServiceMockRecorder {
	return m.recorder
}

// InspectCargo mocks base method.
func (m *MockCargoInspectionService) InspectCargo(arg0 domain.TrackingID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectCargo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InspectCargo indicates an expected call of InspectCargo.
func (mr *MockCargoInspectionServiceMockRecorder) InspectCargo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectCargo", reflect.TypeOf((*MockCargoInspectionService)(nil).InspectCargo), arg0)
}
