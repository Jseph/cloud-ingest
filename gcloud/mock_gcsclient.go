// Code generated by MockGen. DO NOT EDIT.
// Source: gcloud/gcsclient.go

// Package gcloud is a generated GoMock package.
package gcloud

import (
	storage "cloud.google.com/go/storage"
	context "context"
	helpers "github.com/GoogleCloudPlatform/cloud-ingest/helpers"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockGCS is a mock of GCS interface
type MockGCS struct {
	ctrl     *gomock.Controller
	recorder *MockGCSMockRecorder
}

// MockGCSMockRecorder is the mock recorder for MockGCS
type MockGCSMockRecorder struct {
	mock *MockGCS
}

// NewMockGCS creates a new mock instance
func NewMockGCS(ctrl *gomock.Controller) *MockGCS {
	mock := &MockGCS{ctrl: ctrl}
	mock.recorder = &MockGCSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGCS) EXPECT() *MockGCSMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method
func (m *MockGCS) CreateBucket(ctx context.Context, projectId, bucketName string, attrs *storage.BucketAttrs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", ctx, projectId, bucketName, attrs)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockGCSMockRecorder) CreateBucket(ctx, projectId, bucketName, attrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockGCS)(nil).CreateBucket), ctx, projectId, bucketName, attrs)
}

// DeleteBucket mocks base method
func (m *MockGCS) DeleteBucket(ctx context.Context, bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", ctx, bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket
func (mr *MockGCSMockRecorder) DeleteBucket(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockGCS)(nil).DeleteBucket), ctx, bucketName)
}

// DeleteObject mocks base method
func (m *MockGCS) DeleteObject(ctx context.Context, bucketName, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", ctx, bucketName, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockGCSMockRecorder) DeleteObject(ctx, bucketName, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockGCS)(nil).DeleteObject), ctx, bucketName, objectName)
}

// GetAttrs mocks base method
func (m *MockGCS) GetAttrs(ctx context.Context, bucketName, objectName string) (*storage.ObjectAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttrs", ctx, bucketName, objectName)
	ret0, _ := ret[0].(*storage.ObjectAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttrs indicates an expected call of GetAttrs
func (mr *MockGCSMockRecorder) GetAttrs(ctx, bucketName, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttrs", reflect.TypeOf((*MockGCS)(nil).GetAttrs), ctx, bucketName, objectName)
}

// ListObjects mocks base method
func (m *MockGCS) ListObjects(ctx context.Context, bucketName string, query *storage.Query) ObjectIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjects", ctx, bucketName, query)
	ret0, _ := ret[0].(ObjectIterator)
	return ret0
}

// ListObjects indicates an expected call of ListObjects
func (mr *MockGCSMockRecorder) ListObjects(ctx, bucketName, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjects", reflect.TypeOf((*MockGCS)(nil).ListObjects), ctx, bucketName, query)
}

// NewRangeReader mocks base method
func (m *MockGCS) NewRangeReader(ctx context.Context, bucketName, objectName string, offset, length int64) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRangeReader", ctx, bucketName, objectName, offset, length)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewRangeReader indicates an expected call of NewRangeReader
func (mr *MockGCSMockRecorder) NewRangeReader(ctx, bucketName, objectName, offset, length interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRangeReader", reflect.TypeOf((*MockGCS)(nil).NewRangeReader), ctx, bucketName, objectName, offset, length)
}

// NewWriter mocks base method
func (m *MockGCS) NewWriter(ctx context.Context, bucketName, objectName string) helpers.WriteCloserWithError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWriter", ctx, bucketName, objectName)
	ret0, _ := ret[0].(helpers.WriteCloserWithError)
	return ret0
}

// NewWriter indicates an expected call of NewWriter
func (mr *MockGCSMockRecorder) NewWriter(ctx, bucketName, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWriter", reflect.TypeOf((*MockGCS)(nil).NewWriter), ctx, bucketName, objectName)
}

// NewWriterWithCondition mocks base method
func (m *MockGCS) NewWriterWithCondition(ctx context.Context, bucketName, objectName string, cond storage.Conditions) helpers.WriteCloserWithError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWriterWithCondition", ctx, bucketName, objectName, cond)
	ret0, _ := ret[0].(helpers.WriteCloserWithError)
	return ret0
}

// NewWriterWithCondition indicates an expected call of NewWriterWithCondition
func (mr *MockGCSMockRecorder) NewWriterWithCondition(ctx, bucketName, objectName, cond interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWriterWithCondition", reflect.TypeOf((*MockGCS)(nil).NewWriterWithCondition), ctx, bucketName, objectName, cond)
}

// MockObjectIterator is a mock of ObjectIterator interface
type MockObjectIterator struct {
	ctrl     *gomock.Controller
	recorder *MockObjectIteratorMockRecorder
}

// MockObjectIteratorMockRecorder is the mock recorder for MockObjectIterator
type MockObjectIteratorMockRecorder struct {
	mock *MockObjectIterator
}

// NewMockObjectIterator creates a new mock instance
func NewMockObjectIterator(ctrl *gomock.Controller) *MockObjectIterator {
	mock := &MockObjectIterator{ctrl: ctrl}
	mock.recorder = &MockObjectIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObjectIterator) EXPECT() *MockObjectIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockObjectIterator) Next() (*storage.ObjectAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*storage.ObjectAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockObjectIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockObjectIterator)(nil).Next))
}
