// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: AddonsService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	go_scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAddonsService is a mock of AddonsService interface
type MockAddonsService struct {
	ctrl     *gomock.Controller
	recorder *MockAddonsServiceMockRecorder
}

// MockAddonsServiceMockRecorder is the mock recorder for MockAddonsService
type MockAddonsServiceMockRecorder struct {
	mock *MockAddonsService
}

// NewMockAddonsService creates a new mock instance
func NewMockAddonsService(ctrl *gomock.Controller) *MockAddonsService {
	mock := &MockAddonsService{ctrl: ctrl}
	mock.recorder = &MockAddonsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAddonsService) EXPECT() *MockAddonsServiceMockRecorder {
	return m.recorder
}

// AddonDestroy mocks base method
func (m *MockAddonsService) AddonDestroy(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AddonDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddonDestroy indicates an expected call of AddonDestroy
func (mr *MockAddonsServiceMockRecorder) AddonDestroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonDestroy", reflect.TypeOf((*MockAddonsService)(nil).AddonDestroy), arg0, arg1)
}

// AddonProvision mocks base method
func (m *MockAddonsService) AddonProvision(arg0, arg1, arg2 string) (go_scalingo.AddonRes, error) {
	ret := m.ctrl.Call(m, "AddonProvision", arg0, arg1, arg2)
	ret0, _ := ret[0].(go_scalingo.AddonRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonProvision indicates an expected call of AddonProvision
func (mr *MockAddonsServiceMockRecorder) AddonProvision(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonProvision", reflect.TypeOf((*MockAddonsService)(nil).AddonProvision), arg0, arg1, arg2)
}

// AddonUpgrade mocks base method
func (m *MockAddonsService) AddonUpgrade(arg0, arg1, arg2 string) (go_scalingo.AddonRes, error) {
	ret := m.ctrl.Call(m, "AddonUpgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(go_scalingo.AddonRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonUpgrade indicates an expected call of AddonUpgrade
func (mr *MockAddonsServiceMockRecorder) AddonUpgrade(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonUpgrade", reflect.TypeOf((*MockAddonsService)(nil).AddonUpgrade), arg0, arg1, arg2)
}

// AddonsList mocks base method
func (m *MockAddonsService) AddonsList(arg0 string) ([]*go_scalingo.Addon, error) {
	ret := m.ctrl.Call(m, "AddonsList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Addon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonsList indicates an expected call of AddonsList
func (mr *MockAddonsServiceMockRecorder) AddonsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonsList", reflect.TypeOf((*MockAddonsService)(nil).AddonsList), arg0)
}
