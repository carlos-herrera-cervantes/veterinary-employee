// Code generated by MockGen. DO NOT EDIT.
// Source: veterinary-employee/services (interfaces: IAmazonS3)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go/service/s3"
	gomock "github.com/golang/mock/gomock"
)

// MockIAmazonS3 is a mock of IAmazonS3 interface.
type MockIAmazonS3 struct {
	ctrl     *gomock.Controller
	recorder *MockIAmazonS3MockRecorder
}

// MockIAmazonS3MockRecorder is the mock recorder for MockIAmazonS3.
type MockIAmazonS3MockRecorder struct {
	mock *MockIAmazonS3
}

// NewMockIAmazonS3 creates a new mock instance.
func NewMockIAmazonS3(ctrl *gomock.Controller) *MockIAmazonS3 {
	mock := &MockIAmazonS3{ctrl: ctrl}
	mock.recorder = &MockIAmazonS3MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAmazonS3) EXPECT() *MockIAmazonS3MockRecorder {
	return m.recorder
}

// DeleteObject mocks base method.
func (m *MockIAmazonS3) DeleteObject(arg0 *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0)
	ret0, _ := ret[0].(*s3.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockIAmazonS3MockRecorder) DeleteObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockIAmazonS3)(nil).DeleteObject), arg0)
}

// PutObject mocks base method.
func (m *MockIAmazonS3) PutObject(arg0 *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", arg0)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockIAmazonS3MockRecorder) PutObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockIAmazonS3)(nil).PutObject), arg0)
}
