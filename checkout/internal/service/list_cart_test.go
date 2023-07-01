//go:build unit
// +build unit

package service

import (
	"context"
	"errors"
	"testing"

	wp "route256/libs/worker_pool"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"route256/checkout/internal/models"
)

func Test_cartService_ListCart(t *testing.T) {
	tests := []struct {
		name           string
		user           int64
		repoMock       func(m *MockRepository)
		psClientMock   func(m *MockPSClient)
		wantTotalPrice uint32
		wantErr        error
	}{
		{
			name: "success",
			user: 1,
			repoMock: func(m *MockRepository) {
				m.EXPECT().GetUserCart(context.Background(), int64(1)).Return([]models.ItemData{
					{SKU: 1, Count: 2},
					{SKU: 2, Count: 3},
				}, nil).Times(1)
			},
			psClientMock: func(m *MockPSClient) {
				m.EXPECT().GetProducts(context.Background(), []models.ItemData{
					{SKU: 1, Count: 2},
					{SKU: 2, Count: 3},
				}).Return([]wp.Result[models.Item]{
					{Value: models.Item{ItemData: models.ItemData{SKU: 1, Count: 2}, ItemInfo: models.ItemInfo{Name: "test", Price: 5}}, Err: nil},
					{Value: models.Item{ItemData: models.ItemData{SKU: 2, Count: 3}, ItemInfo: models.ItemInfo{Name: "test2", Price: 10}}, Err: nil},
				}).Times(1)
			},
			wantTotalPrice: 40,
			wantErr:        nil,
		},
		{
			name: "fail repo error",
			user: 1,
			repoMock: func(m *MockRepository) {
				m.EXPECT().GetUserCart(context.Background(), int64(1)).Return(nil, errors.New("some repo error")).Times(1)
			},
			wantTotalPrice: 0,
			wantErr:        errors.New("some repo error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := NewMockRepository(mockCtrl)
			mockPSClient := NewMockPSClient(mockCtrl)

			s := &cartService{
				repo:     mockRepo,
				psClient: mockPSClient,
			}

			tt.repoMock(mockRepo)
			if tt.psClientMock != nil {
				tt.psClientMock(mockPSClient)
			}

			cartInfo, err := s.ListCart(context.Background(), tt.user)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantTotalPrice, cartInfo.TotalPrice)
			}
		})
	}
}
