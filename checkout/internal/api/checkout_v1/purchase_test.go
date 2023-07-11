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

func TestImplementation_Purchase(t *testing.T) {
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
		OrderID        int64
		mockServiceErr error
		wantErr        bool
		wantCode       codes.Code
	}{
		{
			name: "fail",
			req: &desc.OrderIDRequest{
				User: 1,
			},
			OrderID:        0,
			mockServiceErr: errors.New("failed to purchase"),
			wantErr:        true,
			wantCode:       codes.Internal,
		},
		{
			name: "success",
			req: &desc.OrderIDRequest{
				User: 1,
			},
			OrderID:        1,
			mockServiceErr: nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCartService.EXPECT().
				Purchase(ctx, tt.req.GetUser()).
				Return(tt.OrderID, tt.mockServiceErr).
				Times(1)

			_, err := impl.Purchase(ctx, tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantCode, status.Code(err))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
