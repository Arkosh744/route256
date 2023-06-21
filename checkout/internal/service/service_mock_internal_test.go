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

// AddToCart mocks base method.
func (m *MockRepository) AddToCart(ctx context.Context, user int64, item *models.ItemData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCart", ctx, user, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToCart indicates an expected call of AddToCart.
func (mr *MockRepositoryMockRecorder) AddToCart(ctx, user, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCart", reflect.TypeOf((*MockRepository)(nil).AddToCart), ctx, user, item)
}

// DeleteFromCart mocks base method.
func (m *MockRepository) DeleteFromCart(ctx context.Context, user int64, item *models.ItemData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFromCart", ctx, user, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFromCart indicates an expected call of DeleteFromCart.
func (mr *MockRepositoryMockRecorder) DeleteFromCart(ctx, user, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromCart", reflect.TypeOf((*MockRepository)(nil).DeleteFromCart), ctx, user, item)
}

// DeleteUserCart mocks base method.
func (m *MockRepository) DeleteUserCart(ctx context.Context, user int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserCart", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserCart indicates an expected call of DeleteUserCart.
func (mr *MockRepositoryMockRecorder) DeleteUserCart(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserCart", reflect.TypeOf((*MockRepository)(nil).DeleteUserCart), ctx, user)
}

// GetCount mocks base method.
func (m *MockRepository) GetCount(ctx context.Context, user int64, sku uint32) (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", ctx, user, sku)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockRepositoryMockRecorder) GetCount(ctx, user, sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockRepository)(nil).GetCount), ctx, user, sku)
}

// GetUserCart mocks base method.
func (m *MockRepository) GetUserCart(ctx context.Context, user int64) ([]models.ItemData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCart", ctx, user)
	ret0, _ := ret[0].([]models.ItemData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCart indicates an expected call of GetUserCart.
func (mr *MockRepositoryMockRecorder) GetUserCart(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCart", reflect.TypeOf((*MockRepository)(nil).GetUserCart), ctx, user)
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
func (m *MockLomsClient) CreateOrder(ctx context.Context, user int64, items []models.ItemData) (int64, error) {
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

// GetProducts mocks base method.
func (m *MockPSClient) GetProducts(ctx context.Context, userItems []models.ItemData) ItemsResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", ctx, userItems)
	ret0, _ := ret[0].(ItemsResult)
	return ret0
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockPSClientMockRecorder) GetProducts(ctx, userItems interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockPSClient)(nil).GetProducts), ctx, userItems)
}
