// Code generated by MockGen. DO NOT EDIT.
// Source: src/repositories/auth/auth.go

// Package mock_premium_feature is a generated GoMock package
package mock_premium_feature

import (
	"context"
	"reflect"
	"DatingApp/src/filter"
	"DatingApp/src/models"

	"github.com/golang/mock/gomock"
)

type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

func (m *MockInterface) Create(ctx context.Context, input models.Query[models.PremiumFeatureInput]) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockInterfaceMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockInterface)(nil).Create), ctx, input)
}

func (m *MockInterface) Get(ctx context.Context, paging filter.Paging[filter.PremiumFeatureFilter]) ([]models.PremiumFeature, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, paging)
	ret0, _ := ret[0].([]models.PremiumFeature)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockInterfaceMockRecorder) Get(ctx, paging interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, paging)
}

func (m *MockInterface) Update(ctx context.Context, input models.Query[models.PremiumFeatureInput], id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, input, id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockInterfaceMockRecorder) Update(ctx, input, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockInterface)(nil).Update), ctx, input, id)
}
