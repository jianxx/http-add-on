// /*
// Copyright 2023 The KEDA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: operator/generated/informers/externalversions/internalinterfaces/factory_interfaces.go
//
// Generated by this command:
//
//	mockgen -copyright_file=hack/boilerplate.go.txt -destination=operator/generated/informers/externalversions/internalinterfaces/mock/factory_interfaces.go -package=mock -source=operator/generated/informers/externalversions/internalinterfaces/factory_interfaces.go
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	internalinterfaces "github.com/kedacore/http-add-on/operator/generated/informers/externalversions/internalinterfaces"
	gomock "go.uber.org/mock/gomock"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cache "k8s.io/client-go/tools/cache"
)

// MockSharedInformerFactory is a mock of SharedInformerFactory interface.
type MockSharedInformerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSharedInformerFactoryMockRecorder
	isgomock struct{}
}

// MockSharedInformerFactoryMockRecorder is the mock recorder for MockSharedInformerFactory.
type MockSharedInformerFactoryMockRecorder struct {
	mock *MockSharedInformerFactory
}

// NewMockSharedInformerFactory creates a new mock instance.
func NewMockSharedInformerFactory(ctrl *gomock.Controller) *MockSharedInformerFactory {
	mock := &MockSharedInformerFactory{ctrl: ctrl}
	mock.recorder = &MockSharedInformerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharedInformerFactory) EXPECT() *MockSharedInformerFactoryMockRecorder {
	return m.recorder
}

// InformerFor mocks base method.
func (m *MockSharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InformerFor", obj, newFunc)
	ret0, _ := ret[0].(cache.SharedIndexInformer)
	return ret0
}

// InformerFor indicates an expected call of InformerFor.
func (mr *MockSharedInformerFactoryMockRecorder) InformerFor(obj, newFunc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InformerFor", reflect.TypeOf((*MockSharedInformerFactory)(nil).InformerFor), obj, newFunc)
}

// Start mocks base method.
func (m *MockSharedInformerFactory) Start(stopCh <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", stopCh)
}

// Start indicates an expected call of Start.
func (mr *MockSharedInformerFactoryMockRecorder) Start(stopCh any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSharedInformerFactory)(nil).Start), stopCh)
}
