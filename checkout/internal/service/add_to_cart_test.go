package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"route256/checkout/internal/models"
)

func Test_cartService_AddToCart(t *testing.T) {
	type args struct {
		user  int64
		sku   uint32
		count uint16
	}

	tests := []struct {
		name           string
		args           args
		repoMock       func(mockRepo *MockRepository)
		lomsClientMock func(m *MockLomsClient)
		wantErr        error
	}{
		{
			name: "fail insufficient stock",
			args: args{
				user:  1,
				sku:   1,
				count: 1,
			},
			repoMock: func(m *MockRepository) {
				m.EXPECT().GetCount(context.Background(), int64(1), uint32(1)).Return(uint16(0), nil).Times(1)
			},
			lomsClientMock: func(m *MockLomsClient) {
				m.EXPECT().Stocks(context.Background(), uint32(1)).Return([]*models.Stock{}, nil).Times(1)
			},
			wantErr: ErrStockInsufficient,
		},
		{
			name: "fail stocks error",
			args: args{
				user:  1,
				sku:   1,
				count: 1,
			},
			repoMock: func(m *MockRepository) {},
			lomsClientMock: func(m *MockLomsClient) {
				m.EXPECT().Stocks(context.Background(), uint32(1)).Return([]*models.Stock{}, errors.New("test")).Times(1)
			},
			wantErr: errors.New("get stocks: test"),
		},
		{
			name: "success",
			args: args{
				user:  1,
				sku:   1,
				count: 1,
			},
			repoMock: func(m *MockRepository) {
				m.EXPECT().GetCount(context.Background(), int64(1), uint32(1)).Return(uint16(0), nil).Times(1)
				m.EXPECT().AddToCart(context.Background(), int64(1), &models.ItemData{SKU: uint32(1), Count: uint16(1)}).Return(nil).Times(1)
			},
			lomsClientMock: func(m *MockLomsClient) {
				m.EXPECT().Stocks(context.Background(), uint32(1)).Return([]*models.Stock{
					{Count: 1},
				}, nil).Times(1)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := NewMockRepository(mockCtrl)
			mockLomsClient := NewMockLomsClient(mockCtrl)

			s := &cartService{
				repo:       mockRepo,
				lomsClient: mockLomsClient,
			}

			tt.repoMock(mockRepo)
			tt.lomsClientMock(mockLomsClient)

			err := s.AddToCart(context.Background(), tt.args.user, tt.args.sku, tt.args.count)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
