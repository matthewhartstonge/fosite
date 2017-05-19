// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: AccessTokenStrategy)

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/MatthewHartstonge/fosite"
)

// Mock of AccessTokenStrategy interface
type MockAccessTokenStrategy struct {
	ctrl     *gomock.Controller
	recorder *_MockAccessTokenStrategyRecorder
}

// Recorder for MockAccessTokenStrategy (not exported)
type _MockAccessTokenStrategyRecorder struct {
	mock *MockAccessTokenStrategy
}

func NewMockAccessTokenStrategy(ctrl *gomock.Controller) *MockAccessTokenStrategy {
	mock := &MockAccessTokenStrategy{ctrl: ctrl}
	mock.recorder = &_MockAccessTokenStrategyRecorder{mock}
	return mock
}

func (_m *MockAccessTokenStrategy) EXPECT() *_MockAccessTokenStrategyRecorder {
	return _m.recorder
}

func (_m *MockAccessTokenStrategy) AccessTokenSignature(_param0 string) string {
	ret := _m.ctrl.Call(_m, "AccessTokenSignature", _param0)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockAccessTokenStrategyRecorder) AccessTokenSignature(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AccessTokenSignature", arg0)
}

func (_m *MockAccessTokenStrategy) GenerateAccessToken(_param0 context.Context, _param1 fosite.Requester) (string, string, error) {
	ret := _m.ctrl.Call(_m, "GenerateAccessToken", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAccessTokenStrategyRecorder) GenerateAccessToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateAccessToken", arg0, arg1)
}

func (_m *MockAccessTokenStrategy) ValidateAccessToken(_param0 context.Context, _param1 fosite.Requester, _param2 string) error {
	ret := _m.ctrl.Call(_m, "ValidateAccessToken", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessTokenStrategyRecorder) ValidateAccessToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateAccessToken", arg0, arg1, arg2)
}
