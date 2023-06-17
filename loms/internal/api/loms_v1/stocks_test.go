package loms_v1

import (
	"context"
	"errors"
	"testing"

	"route256/loms/internal/converter"
	"route256/loms/internal/models"
	desc "route256/pkg/loms_v1"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_Stocks(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := NewMockService(mockCtrl)

	impl := &Implementation{
		lomsService: mockService,
	}

	tests := []struct {
		name           string
		req            *desc.StocksRequest
		res            []models.StockItem
		mockServiceErr error
		wantErr        bool
		wantCode       codes.Code
	}{
		{
			name:           "fail",
			req:            &desc.StocksRequest{Sku: 1},
			mockServiceErr: errors.New("error get stocks"),
			wantErr:        true,
			wantCode:       codes.Internal,
		},
		{
			name: "success",
			req:  &desc.StocksRequest{Sku: 1},
			res: []models.StockItem{
				{
					WarehouseID: 1,
					Count:       1,
				},
			},
			mockServiceErr: nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.EXPECT().Stocks(ctx, tt.req.GetSku()).
				Return(tt.res, tt.mockServiceErr).
				Times(1)

			res, err := impl.Stocks(ctx, tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantCode, status.Code(err))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, converter.ToStocksDesc(tt.res), res.GetStocks())
			}
		})
	}
}
