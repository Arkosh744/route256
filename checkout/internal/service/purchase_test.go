//go:build unit
// +build unit

package service

import (
	"context"
	"errors"
	"testing"

	"route256/checkout/internal/models"

	"github.com/golang/mock/gomock"
)

func Test_Purchase(t *testing.T) {
	tests := []struct {
		name     string
		user     int64
		mockFunc func(mockRepo *MockRepository, mockLomsClient *MockLomsClient)
		wantErr  bool
	}{
		{
			name: "success",
			user: 1,
			mockFunc: func(mockRepo *MockRepository, mockLomsClient *MockLomsClient) {
				mockRepo.EXPECT().GetUserCart(gomock.Any(), gomock.Any()).Return([]models.ItemData{{SKU: 1, Count: 2}}, nil)
				mockLomsClient.EXPECT().CreateOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1234), nil)
				mockRepo.EXPECT().DeleteUserCart(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "fail empty cart",
			user: 1,
			mockFunc: func(mockRepo *MockRepository, mockLomsClient *MockLomsClient) {
				mockRepo.EXPECT().GetUserCart(gomock.Any(), gomock.Any()).Return([]models.ItemData{}, nil)
			},
			wantErr: true,
		},
		{
			name: "fail repo error",
			user: 1,
			mockFunc: func(mockRepo *MockRepository, mockLomsClient *MockLomsClient) {
				mockRepo.EXPECT().GetUserCart(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "fail order creation error",
			user: 1,
			mockFunc: func(mockRepo *MockRepository, mockLomsClient *MockLomsClient) {
				mockRepo.EXPECT().GetUserCart(gomock.Any(), gomock.Any()).Return([]models.ItemData{{SKU: 1, Count: 2}}, nil)
				mockLomsClient.EXPECT().CreateOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "fail deletion error",
			user: 1,
			mockFunc: func(mockRepo *MockRepository, mockLomsClient *MockLomsClient) {
				mockRepo.EXPECT().GetUserCart(gomock.Any(), gomock.Any()).Return([]models.ItemData{{SKU: 1, Count: 2}}, nil)
				mockLomsClient.EXPECT().CreateOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1234), nil)
				mockRepo.EXPECT().DeleteUserCart(gomock.Any(), gomock.Any()).Return(errors.New("some error"))
			},
			wantErr: true,
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

			tt.mockFunc(mockRepo, mockLomsClient)

			_, err := s.Purchase(context.Background(), tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Purchase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
