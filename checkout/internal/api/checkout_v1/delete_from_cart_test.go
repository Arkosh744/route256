//go:build unit
// +build unit

package checkout_v1

import (
	"context"
	"errors"
	"testing"

	desc "route256/pkg/checkout_v1"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_DeleteFromCart(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCartService := NewMockService(mockCtrl)

	impl := &Implementation{
		cartService: mockCartService,
	}

	tests := []struct {
		name           string
		req            *desc.CartRequest
		mockServiceErr error
		wantErr        bool
		wantCode       codes.Code
	}{
		{
			name: "fail",
			req: &desc.CartRequest{
				User:  1,
				Sku:   1,
				Count: 1,
			},
			mockServiceErr: errors.New("failed to delete from cart"),
			wantErr:        true,
			wantCode:       codes.Internal,
		},
		{
			name: "success",
			req: &desc.CartRequest{
				User:  1,
				Sku:   1,
				Count: 1,
			},
			mockServiceErr: nil,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCartService.EXPECT().
				DeleteFromCart(ctx, tt.req.GetUser(), tt.req.GetSku(), uint16(tt.req.GetCount())).
				Return(tt.mockServiceErr).
				Times(1)

			_, err := impl.DeleteFromCart(ctx, tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantCode, status.Code(err))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
