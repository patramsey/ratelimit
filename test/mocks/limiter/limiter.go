// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/envoyproxy/ratelimit/src/limiter (interfaces: RateLimitCache)

// Package mock_limiter is a generated GoMock package.
package mock_limiter

import (
	context "context"
	config "github.com/envoyproxy/ratelimit/src/config"
	gomock "github.com/golang/mock/gomock"
	envoy_service_ratelimit_v3 "github.com/patramsey/go-control-plane/envoy/service/ratelimit/v3"
	reflect "reflect"
)

// MockRateLimitCache is a mock of RateLimitCache interface
type MockRateLimitCache struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimitCacheMockRecorder
}

// MockRateLimitCacheMockRecorder is the mock recorder for MockRateLimitCache
type MockRateLimitCacheMockRecorder struct {
	mock *MockRateLimitCache
}

// NewMockRateLimitCache creates a new mock instance
func NewMockRateLimitCache(ctrl *gomock.Controller) *MockRateLimitCache {
	mock := &MockRateLimitCache{ctrl: ctrl}
	mock.recorder = &MockRateLimitCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRateLimitCache) EXPECT() *MockRateLimitCacheMockRecorder {
	return m.recorder
}

// DoLimit mocks base method
func (m *MockRateLimitCache) DoReset(arg0 context.Context, arg1 *envoy_service_ratelimit_v3.RateLimitRequest, arg2 []*config.RateLimit) []*envoy_service_ratelimit_v3.RateLimitResponse_DescriptorStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoReset", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*envoy_service_ratelimit_v3.RateLimitResponse_DescriptorStatus)
	return ret0
}

// DoLimit indicates an expected call of DoLimit
func (mr *MockRateLimitCacheMockRecorder) DoReset(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoReset", reflect.TypeOf((*MockRateLimitCache)(nil).DoReset), arg0, arg1, arg2)
}

// DoLimit mocks base method
func (m *MockRateLimitCache) DoLimit(arg0 context.Context, arg1 *envoy_service_ratelimit_v3.RateLimitRequest, arg2 []*config.RateLimit) []*envoy_service_ratelimit_v3.RateLimitResponse_DescriptorStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoLimit", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*envoy_service_ratelimit_v3.RateLimitResponse_DescriptorStatus)
	return ret0
}

// DoLimit indicates an expected call of DoLimit
func (mr *MockRateLimitCacheMockRecorder) DoLimit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoLimit", reflect.TypeOf((*MockRateLimitCache)(nil).DoLimit), arg0, arg1, arg2)
}

// Flush mocks base method
func (m *MockRateLimitCache) Flush() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush")
}

// Flush indicates an expected call of Flush
func (mr *MockRateLimitCacheMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockRateLimitCache)(nil).Flush))
}
