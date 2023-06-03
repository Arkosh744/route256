package service

import (
	"context"
	"testing"

	"route256/checkout/internal/log"
	"route256/checkout/internal/models"

	"github.com/golang/mock/gomock"
)

func Test_cartService_AddToCart(t *testing.T) {
	ctx := context.Background()
	err := log.InitLogger(ctx)
	if err != nil {
		t.Fatalf("error initializing logger: %v", err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := NewMockRepository(mockCtrl)
	mockLomsClient := NewMockLomsClient(mockCtrl)
	mockPSClient := NewMockPSClient(mockCtrl)

	s := &cartService{
		repo:       mockRepo,
		lomsClient: mockLomsClient,
		psClient:   mockPSClient,
	}

	type args struct {
		user  int64
		sku   uint32
		count uint16
	}

	tests := []struct {
		name          string
		args          args
		mockStocks    []*models.Stock
		mockStocksErr error
		wantErr       bool
	}{
		{
			name: "fail insufficient stock",
			args: args{
				user:  1,
				sku:   1,
				count: 1,
			},
			mockStocks:    []*models.Stock{},
			mockStocksErr: nil,
			wantErr:       true,
		},
		{
			name: "success",
			args: args{
				user:  1,
				sku:   1,
				count: 1,
			},
			mockStocks: []*models.Stock{
				{Count: 1},
			},
			mockStocksErr: nil,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLomsClient.EXPECT().Stocks(ctx, tt.args.sku).Return(tt.mockStocks, tt.mockStocksErr).Times(1)

			if err := s.AddToCart(ctx, tt.args.user, tt.args.sku, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("AddToCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
