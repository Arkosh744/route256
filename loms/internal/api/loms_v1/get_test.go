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

func TestImplementation_ListOrder(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := NewMockService(mockCtrl)

	impl := &Implementation{
		lomsService: mockService,
	}

	tests := []struct {
		name           string
		req            *desc.OrderIDRequest
		res            *models.Order
		mockServiceErr error
		wantErr        bool
		wantCode       codes.Code
	}{
		{
			name: "fail",
			req: &desc.OrderIDRequest{
				OrderId: 1,
			},
			mockServiceErr: errors.New("error get order"),
			wantErr:        true,
			wantCode:       codes.Internal,
		},
		{
			name: "success",
			req: &desc.OrderIDRequest{
				OrderId: 1,
			},
			res: &models.Order{
				Status: models.OrderStatusPaid,
				User:   1,
				Items: []models.Item{
					{
						SKU:   1,
						Count: 1,
					},
				},
			},
			mockServiceErr: nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.EXPECT().Get(ctx, tt.req.GetOrderId()).
				Return(tt.mockServiceErr).
				Times(1)

			res, err := impl.ListOrder(ctx, tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantCode, status.Code(err))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.res.Status, res.GetStatus())
				assert.Equal(t, tt.res.User, res.GetUser())
				assert.Equal(t, converter.ToItemsDesc(tt.res.Items), res.GetItems())
			}
		})
	}
}
