package loms_v1

import (
	"context"
	"errors"
	"testing"

	"route256/loms/internal/converter"
	desc "route256/pkg/loms_v1"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_CreateOrder(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := NewMockService(mockCtrl)

	impl := &Implementation{
		lomsService: mockService,
	}

	tests := []struct {
		name           string
		req            *desc.CreateOrderRequest
		res            int64
		mockServiceErr error
		wantErr        bool
		wantCode       codes.Code
	}{
		{
			name: "fail",
			req: &desc.CreateOrderRequest{
				User: 1,
				Items: []*desc.Item{
					{
						Sku:   1,
						Count: 1,
					},
				},
			},
			mockServiceErr: errors.New("error to create order"),
			wantErr:        true,
			wantCode:       codes.Internal,
		},
		{
			name: "success",
			req: &desc.CreateOrderRequest{
				User: 1,
				Items: []*desc.Item{
					{
						Sku:   1,
						Count: 1,
					},
				},
			},
			res:            1,
			mockServiceErr: nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.EXPECT().Create(ctx, tt.req.GetUser(), converter.ToItems(tt.req.GetItems())).
				Return(tt.res, tt.mockServiceErr).
				Times(1)

			res, err := impl.CreateOrder(ctx, tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantCode, status.Code(err))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.res, res.GetOrderId())
			}
		})
	}
}
