package loms_v1

import (
	"context"
	"errors"
	"testing"

	desc "route256/pkg/loms_v1"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_Paid(t *testing.T) {
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
		mockServiceErr error
		wantErr        bool
		wantCode       codes.Code
	}{
		{
			name: "fail",
			req: &desc.OrderIDRequest{
				OrderId: 1,
			},
			mockServiceErr: errors.New("error set order paid"),
			wantErr:        true,
			wantCode:       codes.Internal,
		},
		{
			name: "success",
			req: &desc.OrderIDRequest{
				OrderId: 1,
			},
			mockServiceErr: nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.EXPECT().Paid(ctx, tt.req.GetOrderId()).
				Return(tt.mockServiceErr).
				Times(1)

			_, err := impl.OrderPaid(ctx, tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantCode, status.Code(err))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
