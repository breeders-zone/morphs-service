// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repositories/repositories.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"

	domain "github.com/breeders-zone/morphs-service/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockGenes is a mock of Genes interface.
type MockGenes struct {
	ctrl     *gomock.Controller
	recorder *MockGenesMockRecorder
}

// MockGenesMockRecorder is the mock recorder for MockGenes.
type MockGenesMockRecorder struct {
	mock *MockGenes
}

// NewMockGenes creates a new mock instance.
func NewMockGenes(ctrl *gomock.Controller) *MockGenes {
	mock := &MockGenes{ctrl: ctrl}
	mock.recorder = &MockGenesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenes) EXPECT() *MockGenesMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockGenes) Create(g *domain.Gene) (*domain.Gene, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", g)
	ret0, _ := ret[0].(*domain.Gene)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockGenesMockRecorder) Create(g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGenes)(nil).Create), g)
}

// Delete mocks base method.
func (m *MockGenes) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockGenesMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGenes)(nil).Delete), id)
}

// GetById mocks base method.
func (m *MockGenes) GetById(id int) (*domain.Gene, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*domain.Gene)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockGenesMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockGenes)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockGenes) Update(g *domain.Gene) (*domain.Gene, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", g)
	ret0, _ := ret[0].(*domain.Gene)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockGenesMockRecorder) Update(g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGenes)(nil).Update), g)
}