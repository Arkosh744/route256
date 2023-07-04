package service

import (
	"context"
	"errors"
	"testing"

	"route256/checkout/internal/models"

	"github.com/golang/mock/gomock"
)

func Test_DeleteFromCart(t *testing.T) {
	type args struct {
		user  int64
		sku   uint32
		count uint16
	}

	tests := []struct {
		name     string
		args     args
		mockFunc func(mockRepo *MockRepository)
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				user:  1,
				sku:   1,
				count: 2,
			},
			mockFunc: func(mockRepo *MockRepository) {
				mockRepo.EXPECT().DeleteFromCart(gomock.Any(), gomock.Any(), &models.ItemData{SKU: 1, Count: 2}).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "fail repo error",
			args: args{
				user:  1,
				sku:   1,
				count: 2,
			},
			mockFunc: func(mockRepo *MockRepository) {
				mockRepo.EXPECT().DeleteFromCart(gomock.Any(), gomock.Any(), &models.ItemData{SKU: 1, Count: 2}).Return(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := NewMockRepository(mockCtrl)

			s := &cartService{
				repo: mockRepo,
			}

			tt.mockFunc(mockRepo)

			err := s.DeleteFromCart(context.Background(), tt.args.user, tt.args.sku, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteFromCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
