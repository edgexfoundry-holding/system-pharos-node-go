/*******************************************************************************
 * Copyright 2017 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/
// Code generated by MockGen. DO NOT EDIT.
// Source: src/controller/monitoring/resource/resource.go

// Package mock_resource is a generated GoMock package.
package mock_resource

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCommand is a mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// GetHostResourceInfo mocks base method
func (m *MockCommand) GetHostResourceInfo() (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "GetHostResourceInfo")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostResourceInfo indicates an expected call of GetHostResourceInfo
func (mr *MockCommandMockRecorder) GetHostResourceInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostResourceInfo", reflect.TypeOf((*MockCommand)(nil).GetHostResourceInfo))
}

// GetAppResourceInfo mocks base method
func (m *MockCommand) GetAppResourceInfo(appId string) (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "GetAppResourceInfo", appId)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppResourceInfo indicates an expected call of GetAppResourceInfo
func (mr *MockCommandMockRecorder) GetAppResourceInfo(appId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppResourceInfo", reflect.TypeOf((*MockCommand)(nil).GetAppResourceInfo), appId)
}