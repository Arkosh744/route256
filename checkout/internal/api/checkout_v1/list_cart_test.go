package checkout_v1

import (
	"context"
	"errors"
	"testing"

	"route256/checkout/internal/converter"
	"route256/checkout/internal/models"
	desc "route256/pkg/checkout_v1"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_ListCart(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCartService := NewMockService(mockCtrl)

	impl := &Implementation{
		cartService: mockCartService,
	}

	tests := []struct {
		name           string
		req            *desc.OrderIDRequest
		resInfo        *models.CartInfo
		mockServiceErr error
		wantErr        bool
		wantCode       codes.Code
	}{
		{
			name: "fail",
			req: &desc.OrderIDRequest{
				User: 1,
			},
			resInfo:        &models.CartInfo{},
			mockServiceErr: errors.New("failed to purchase"),
			wantErr:        true,
			wantCode:       codes.Internal,
		},
		{
			name: "success",
			req: &desc.OrderIDRequest{
				User: 1,
			},
			resInfo: &models.CartInfo{
				Items: []models.ItemInfo{
					{
						ItemBase: models.ItemBase{
							Price: 1,
							Name:  "test",
						},
						ItemData: models.ItemData{
							SKU:   1,
							Count: 1,
						},
					},
					{
						ItemBase: models.ItemBase{
							Price: 2,
							Name:  "test2",
						},
						ItemData: models.ItemData{
							SKU:   2,
							Count: 1,
						}},
				},
				TotalPrice: 3,
			},
			mockServiceErr: nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCartService.EXPECT().
				ListCart(ctx, tt.req.GetUser()).
				Return(tt.resInfo, tt.mockServiceErr).
				Times(1)

			res, err := impl.ListCart(ctx, tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantCode, status.Code(err))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, converter.ListToDesc(tt.resInfo), res)
			}
		})
	}
}
