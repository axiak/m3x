// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3x/ident (interfaces: ID,TagIterator)

// Package ident is a generated GoMock package.
package ident

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockID is a mock of ID interface
type MockID struct {
	ctrl     *gomock.Controller
	recorder *MockIDMockRecorder
}

// MockIDMockRecorder is the mock recorder for MockID
type MockIDMockRecorder struct {
	mock *MockID
}

// NewMockID creates a new mock instance
func NewMockID(ctrl *gomock.Controller) *MockID {
	mock := &MockID{ctrl: ctrl}
	mock.recorder = &MockIDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockID) EXPECT() *MockIDMockRecorder {
	return m.recorder
}

// Bytes mocks base method
func (m *MockID) Bytes() []byte {
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes
func (mr *MockIDMockRecorder) Bytes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockID)(nil).Bytes))
}

// Equal mocks base method
func (m *MockID) Equal(arg0 ID) bool {
	ret := m.ctrl.Call(m, "Equal", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockIDMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockID)(nil).Equal), arg0)
}

// Finalize mocks base method
func (m *MockID) Finalize() {
	m.ctrl.Call(m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (mr *MockIDMockRecorder) Finalize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockID)(nil).Finalize))
}

// IsNoFinalize mocks base method
func (m *MockID) IsNoFinalize() bool {
	ret := m.ctrl.Call(m, "IsNoFinalize")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoFinalize indicates an expected call of IsNoFinalize
func (mr *MockIDMockRecorder) IsNoFinalize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoFinalize", reflect.TypeOf((*MockID)(nil).IsNoFinalize))
}

// NoFinalize mocks base method
func (m *MockID) NoFinalize() {
	m.ctrl.Call(m, "NoFinalize")
}

// NoFinalize indicates an expected call of NoFinalize
func (mr *MockIDMockRecorder) NoFinalize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoFinalize", reflect.TypeOf((*MockID)(nil).NoFinalize))
}

// String mocks base method
func (m *MockID) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockIDMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockID)(nil).String))
}

// MockTagIterator is a mock of TagIterator interface
type MockTagIterator struct {
	ctrl     *gomock.Controller
	recorder *MockTagIteratorMockRecorder
}

// MockTagIteratorMockRecorder is the mock recorder for MockTagIterator
type MockTagIteratorMockRecorder struct {
	mock *MockTagIterator
}

// NewMockTagIterator creates a new mock instance
func NewMockTagIterator(ctrl *gomock.Controller) *MockTagIterator {
	mock := &MockTagIterator{ctrl: ctrl}
	mock.recorder = &MockTagIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagIterator) EXPECT() *MockTagIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockTagIterator) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockTagIteratorMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTagIterator)(nil).Close))
}

// Current mocks base method
func (m *MockTagIterator) Current() Tag {
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(Tag)
	return ret0
}

// Current indicates an expected call of Current
func (mr *MockTagIteratorMockRecorder) Current() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockTagIterator)(nil).Current))
}

// CurrentIndex mocks base method
func (m *MockTagIterator) CurrentIndex() int {
	ret := m.ctrl.Call(m, "CurrentIndex")
	ret0, _ := ret[0].(int)
	return ret0
}

// CurrentIndex indicates an expected call of CurrentIndex
func (mr *MockTagIteratorMockRecorder) CurrentIndex() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentIndex", reflect.TypeOf((*MockTagIterator)(nil).CurrentIndex))
}

// Duplicate mocks base method
func (m *MockTagIterator) Duplicate() TagIterator {
	ret := m.ctrl.Call(m, "Duplicate")
	ret0, _ := ret[0].(TagIterator)
	return ret0
}

// Duplicate indicates an expected call of Duplicate
func (mr *MockTagIteratorMockRecorder) Duplicate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Duplicate", reflect.TypeOf((*MockTagIterator)(nil).Duplicate))
}

// Err mocks base method
func (m *MockTagIterator) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockTagIteratorMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockTagIterator)(nil).Err))
}

// Len mocks base method
func (m *MockTagIterator) Len() int {
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len
func (mr *MockTagIteratorMockRecorder) Len() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockTagIterator)(nil).Len))
}

// Next mocks base method
func (m *MockTagIterator) Next() bool {
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockTagIteratorMockRecorder) Next() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockTagIterator)(nil).Next))
}

// Remaining mocks base method
func (m *MockTagIterator) Remaining() int {
	ret := m.ctrl.Call(m, "Remaining")
	ret0, _ := ret[0].(int)
	return ret0
}

// Remaining indicates an expected call of Remaining
func (mr *MockTagIteratorMockRecorder) Remaining() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remaining", reflect.TypeOf((*MockTagIterator)(nil).Remaining))
}
