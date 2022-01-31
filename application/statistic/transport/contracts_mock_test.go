// Code generated by MockGen. DO NOT EDIT.
// Source: ./contracts.go

// Package transport is a generated GoMock package.
package transport

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/valerii-smirnov/flex-made-query-analizer/application/statistic/services/dto"
)

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// Struct mocks base method.
func (m *MockValidator) Struct(s interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Struct", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// Struct indicates an expected call of Struct.
func (mr *MockValidatorMockRecorder) Struct(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Struct", reflect.TypeOf((*MockValidator)(nil).Struct), s)
}

// MockStatisticService is a mock of StatisticService interface.
type MockStatisticService struct {
	ctrl     *gomock.Controller
	recorder *MockStatisticServiceMockRecorder
}

// MockStatisticServiceMockRecorder is the mock recorder for MockStatisticService.
type MockStatisticServiceMockRecorder struct {
	mock *MockStatisticService
}

// NewMockStatisticService creates a new mock instance.
func NewMockStatisticService(ctrl *gomock.Controller) *MockStatisticService {
	mock := &MockStatisticService{ctrl: ctrl}
	mock.recorder = &MockStatisticServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatisticService) EXPECT() *MockStatisticServiceMockRecorder {
	return m.recorder
}

// GetStatistic mocks base method.
func (m *MockStatisticService) GetStatistic(ctx context.Context, req dto.GetQueriesReq) (dto.QueryStatisticCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatistic", ctx, req)
	ret0, _ := ret[0].(dto.QueryStatisticCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatistic indicates an expected call of GetStatistic.
func (mr *MockStatisticServiceMockRecorder) GetStatistic(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatistic", reflect.TypeOf((*MockStatisticService)(nil).GetStatistic), ctx, req)
}