// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/dinghyfile/builder.go

// Package mock_dinghyfile is a generated GoMock package.
package dinghyfile

import (
	bytes "bytes"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockParser is a mock of Parser interface.
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser.
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance.
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// SetBuilder mocks base method.
func (m *MockParser) SetBuilder(b *PipelineBuilder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBuilder", b)
}

// SetBuilder indicates an expected call of SetBuilder.
func (mr *MockParserMockRecorder) SetBuilder(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBuilder", reflect.TypeOf((*MockParser)(nil).SetBuilder), b)
}

// Parse mocks base method.
func (m *MockParser) Parse(org, repo, path, branch string, vars []VarMap) (*bytes.Buffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", org, repo, path, branch, vars)
	ret0, _ := ret[0].(*bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockParserMockRecorder) Parse(org, repo, path, branch, vars interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockParser)(nil).Parse), org, repo, path, branch, vars)
}

// MockDependencyManager is a mock of DependencyManager interface.
type MockDependencyManager struct {
	ctrl     *gomock.Controller
	recorder *MockDependencyManagerMockRecorder
}

// MockDependencyManagerMockRecorder is the mock recorder for MockDependencyManager.
type MockDependencyManagerMockRecorder struct {
	mock *MockDependencyManager
}

// NewMockDependencyManager creates a new mock instance.
func NewMockDependencyManager(ctrl *gomock.Controller) *MockDependencyManager {
	mock := &MockDependencyManager{ctrl: ctrl}
	mock.recorder = &MockDependencyManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDependencyManager) EXPECT() *MockDependencyManagerMockRecorder {
	return m.recorder
}

// GetRawData mocks base method.
func (m *MockDependencyManager) GetRawData(url string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawData", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawData indicates an expected call of GetRawData.
func (mr *MockDependencyManagerMockRecorder) GetRawData(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawData", reflect.TypeOf((*MockDependencyManager)(nil).GetRawData), url)
}

// SetRawData mocks base method.
func (m *MockDependencyManager) SetRawData(url, rawData string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRawData", url, rawData)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRawData indicates an expected call of SetRawData.
func (mr *MockDependencyManagerMockRecorder) SetRawData(url, rawData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRawData", reflect.TypeOf((*MockDependencyManager)(nil).SetRawData), url, rawData)
}

// SetDeps mocks base method.
func (m *MockDependencyManager) SetDeps(parent string, deps []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeps", parent, deps)
}

// SetDeps indicates an expected call of SetDeps.
func (mr *MockDependencyManagerMockRecorder) SetDeps(parent, deps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeps", reflect.TypeOf((*MockDependencyManager)(nil).SetDeps), parent, deps)
}

// GetRoots mocks base method.
func (m *MockDependencyManager) GetRoots(child string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoots", child)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetRoots indicates an expected call of GetRoots.
func (mr *MockDependencyManagerMockRecorder) GetRoots(child interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoots", reflect.TypeOf((*MockDependencyManager)(nil).GetRoots), child)
}

// MockDownloader is a mock of Downloader interface.
type MockDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockDownloaderMockRecorder
}

// MockDownloaderMockRecorder is the mock recorder for MockDownloader.
type MockDownloaderMockRecorder struct {
	mock *MockDownloader
}

// NewMockDownloader creates a new mock instance.
func NewMockDownloader(ctrl *gomock.Controller) *MockDownloader {
	mock := &MockDownloader{ctrl: ctrl}
	mock.recorder = &MockDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloader) EXPECT() *MockDownloaderMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockDownloader) Download(org, repo, file, branch string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", org, repo, file, branch)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockDownloaderMockRecorder) Download(org, repo, file, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockDownloader)(nil).Download), org, repo, file, branch)
}

// EncodeURL mocks base method.
func (m *MockDownloader) EncodeURL(org, repo, file, branch string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncodeURL", org, repo, file, branch)
	ret0, _ := ret[0].(string)
	return ret0
}

// EncodeURL indicates an expected call of EncodeURL.
func (mr *MockDownloaderMockRecorder) EncodeURL(org, repo, file, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncodeURL", reflect.TypeOf((*MockDownloader)(nil).EncodeURL), org, repo, file, branch)
}

// DecodeURL mocks base method.
func (m *MockDownloader) DecodeURL(url string) (string, string, string, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeURL", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(string)
	return ret0, ret1, ret2, ret3
}

// DecodeURL indicates an expected call of DecodeURL.
func (mr *MockDownloaderMockRecorder) DecodeURL(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeURL", reflect.TypeOf((*MockDownloader)(nil).DecodeURL), url)
}
