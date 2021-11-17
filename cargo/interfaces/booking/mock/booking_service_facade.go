// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fengye87/dddsample-core/cargo/interfaces/booking (interfaces: BookingServiceFacade)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	booking "github.com/fengye87/dddsample-core/cargo/interfaces/booking"
	gomock "github.com/golang/mock/gomock"
)

// MockBookingServiceFacade is a mock of BookingServiceFacade interface.
type MockBookingServiceFacade struct {
	ctrl     *gomock.Controller
	recorder *MockBookingServiceFacadeMockRecorder
}

// MockBookingServiceFacadeMockRecorder is the mock recorder for MockBookingServiceFacade.
type MockBookingServiceFacadeMockRecorder struct {
	mock *MockBookingServiceFacade
}

// NewMockBookingServiceFacade creates a new mock instance.
func NewMockBookingServiceFacade(ctrl *gomock.Controller) *MockBookingServiceFacade {
	mock := &MockBookingServiceFacade{ctrl: ctrl}
	mock.recorder = &MockBookingServiceFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookingServiceFacade) EXPECT() *MockBookingServiceFacadeMockRecorder {
	return m.recorder
}

// AssignCargoToRoute mocks base method.
func (m *MockBookingServiceFacade) AssignCargoToRoute(arg0 string, arg1 *booking.RouteCandidateDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignCargoToRoute", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignCargoToRoute indicates an expected call of AssignCargoToRoute.
func (mr *MockBookingServiceFacadeMockRecorder) AssignCargoToRoute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignCargoToRoute", reflect.TypeOf((*MockBookingServiceFacade)(nil).AssignCargoToRoute), arg0, arg1)
}

// BookNewCargo mocks base method.
func (m *MockBookingServiceFacade) BookNewCargo(arg0, arg1 string, arg2 time.Time) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookNewCargo", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookNewCargo indicates an expected call of BookNewCargo.
func (mr *MockBookingServiceFacadeMockRecorder) BookNewCargo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookNewCargo", reflect.TypeOf((*MockBookingServiceFacade)(nil).BookNewCargo), arg0, arg1, arg2)
}

// ChangeDestination mocks base method.
func (m *MockBookingServiceFacade) ChangeDestination(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeDestination", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeDestination indicates an expected call of ChangeDestination.
func (mr *MockBookingServiceFacadeMockRecorder) ChangeDestination(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeDestination", reflect.TypeOf((*MockBookingServiceFacade)(nil).ChangeDestination), arg0, arg1)
}

// ListAllCargos mocks base method.
func (m *MockBookingServiceFacade) ListAllCargos() ([]booking.CargoRoutingDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllCargos")
	ret0, _ := ret[0].([]booking.CargoRoutingDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllCargos indicates an expected call of ListAllCargos.
func (mr *MockBookingServiceFacadeMockRecorder) ListAllCargos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllCargos", reflect.TypeOf((*MockBookingServiceFacade)(nil).ListAllCargos))
}

// ListShippingLocations mocks base method.
func (m *MockBookingServiceFacade) ListShippingLocations() ([]booking.LocationDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShippingLocations")
	ret0, _ := ret[0].([]booking.LocationDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShippingLocations indicates an expected call of ListShippingLocations.
func (mr *MockBookingServiceFacadeMockRecorder) ListShippingLocations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShippingLocations", reflect.TypeOf((*MockBookingServiceFacade)(nil).ListShippingLocations))
}

// LoadCargoForRouting mocks base method.
func (m *MockBookingServiceFacade) LoadCargoForRouting(arg0 string) (*booking.CargoRoutingDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadCargoForRouting", arg0)
	ret0, _ := ret[0].(*booking.CargoRoutingDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadCargoForRouting indicates an expected call of LoadCargoForRouting.
func (mr *MockBookingServiceFacadeMockRecorder) LoadCargoForRouting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadCargoForRouting", reflect.TypeOf((*MockBookingServiceFacade)(nil).LoadCargoForRouting), arg0)
}

// RequestPossibleRoutesForCargo mocks base method.
func (m *MockBookingServiceFacade) RequestPossibleRoutesForCargo(arg0 string) ([]booking.RouteCandidateDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestPossibleRoutesForCargo", arg0)
	ret0, _ := ret[0].([]booking.RouteCandidateDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestPossibleRoutesForCargo indicates an expected call of RequestPossibleRoutesForCargo.
func (mr *MockBookingServiceFacadeMockRecorder) RequestPossibleRoutesForCargo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestPossibleRoutesForCargo", reflect.TypeOf((*MockBookingServiceFacade)(nil).RequestPossibleRoutesForCargo), arg0)
}
