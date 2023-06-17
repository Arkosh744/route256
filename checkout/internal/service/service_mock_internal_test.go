// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"
	models "route256/checkout/internal/models"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// MockLomsClient is a mock of LomsClient interface.
type MockLomsClient struct {
	ctrl     *gomock.Controller
	recorder *MockLomsClientMockRecorder
}

// MockLomsClientMockRecorder is the mock recorder for MockLomsClient.
type MockLomsClientMockRecorder struct {
	mock *MockLomsClient
}

// NewMockLomsClient creates a new mock instance.
func NewMockLomsClient(ctrl *gomock.Controller) *MockLomsClient {
	mock := &MockLomsClient{ctrl: ctrl}
	mock.recorder = &MockLomsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLomsClient) EXPECT() *MockLomsClientMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockLomsClient) CreateOrder(ctx context.Context, user int64, items []*models.ItemData) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, user, items)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockLomsClientMockRecorder) CreateOrder(ctx, user, items interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockLomsClient)(nil).CreateOrder), ctx, user, items)
}

// Stocks mocks base method.
func (m *MockLomsClient) Stocks(ctx context.Context, sku uint32) ([]*models.Stock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stocks", ctx, sku)
	ret0, _ := ret[0].([]*models.Stock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stocks indicates an expected call of Stocks.
func (mr *MockLomsClientMockRecorder) Stocks(ctx, sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stocks", reflect.TypeOf((*MockLomsClient)(nil).Stocks), ctx, sku)
}

// MockPSClient is a mock of PSClient interface.
type MockPSClient struct {
	ctrl     *gomock.Controller
	recorder *MockPSClientMockRecorder
}

// MockPSClientMockRecorder is the mock recorder for MockPSClient.
type MockPSClientMockRecorder struct {
	mock *MockPSClient
}

// NewMockPSClient creates a new mock instance.
func NewMockPSClient(ctrl *gomock.Controller) *MockPSClient {
	mock := &MockPSClient{ctrl: ctrl}
	mock.recorder = &MockPSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPSClient) EXPECT() *MockPSClientMockRecorder {
	return m.recorder
}

// GetProduct mocks base method.
func (m *MockPSClient) GetProduct(ctx context.Context, sku uint32) (*models.ItemBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, sku)
	ret0, _ := ret[0].(*models.ItemBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockPSClientMockRecorder) GetProduct(ctx, sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockPSClient)(nil).GetProduct), ctx, sku)
}