package service

import (
	"context"
	"database/sql"
	"testing"

	"route256/libs/client/pg"
	"route256/libs/log"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"route256/loms/internal/models"
)

func Test_service_Paid(t *testing.T) {
	ctx := context.Background()
	err := log.InitLogger(ctx, "dev")
	require.NoError(t, err)

	tests := []struct {
		name             string
		orderID          int64
		repoMock         func(m *MockRepository)
		txManagerMock    func(m *pg.MockTxManager)
		kafkaMock        func(m *MockOrderStatusSender)
		orderStorageMock func(m *orderStorage)
		wantErr          error
	}{
		{
			name:    "fail order not found",
			orderID: 1,
			repoMock: func(m *MockRepository) {
				m.EXPECT().GetOrder(ctx, int64(1)).Return(nil, sql.ErrNoRows).Times(1)
			},
			txManagerMock: func(m *pg.MockTxManager) {
				m.EXPECT().RunRepeatableRead(ctx, gomock.Any()).DoAndReturn(
					func(ctx context.Context, f func(context.Context) error) error {
						return f(ctx)
					}).Times(1)
			},
			wantErr: ErrOrderNotFound,
		},
		{
			name:    "fail invalid order status",
			orderID: 1,
			repoMock: func(m *MockRepository) {
				m.EXPECT().GetOrder(ctx, int64(1)).Return(&models.Order{Status: models.OrderStatusFailed}, nil).Times(1)
			},
			txManagerMock: func(m *pg.MockTxManager) {
				m.EXPECT().RunRepeatableRead(ctx, gomock.Any()).DoAndReturn(
					func(ctx context.Context, f func(context.Context) error) error {
						return f(ctx)
					}).Times(1)
			},
			wantErr: ErrInvalidOrderStatus,
		},
		{
			name:    "success",
			orderID: 1,
			repoMock: func(m *MockRepository) {
				m.EXPECT().GetOrder(ctx, int64(1)).Return(&models.Order{Status: models.OrderStatusAwaitingPayment}, nil).Times(1)
				m.EXPECT().DeleteReservation(ctx, int64(1)).Return(nil).Times(1)
				m.EXPECT().UpdateOrderStatus(ctx, int64(1), models.OrderStatusPaid).Return(nil).Times(1)
			},
			txManagerMock: func(m *pg.MockTxManager) {
				m.EXPECT().RunRepeatableRead(ctx, gomock.Any()).DoAndReturn(
					func(ctx context.Context, f func(context.Context) error) error {
						return f(ctx)
					}).Times(1)
			},
			kafkaMock: func(m *MockOrderStatusSender) {
				m.EXPECT().SendOrderStatus(int64(1), models.OrderStatusPaid).Return(nil).Times(1)
			},
			orderStorageMock: func(m *orderStorage) {
				m.deleteFromStorage(1)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := NewMockRepository(mockCtrl)
			mockTxManager := pg.NewMockTxManager(mockCtrl)
			mockKafka := NewMockOrderStatusSender(mockCtrl)

			s := &service{
				repo:      mockRepo,
				txManager: mockTxManager,
				kafka:     mockKafka,
				storage:   orderStorage{storage: make(map[int64]*orderStatus)},
			}

			if tt.repoMock != nil {
				tt.repoMock(mockRepo)
			}
			if tt.txManagerMock != nil {
				tt.txManagerMock(mockTxManager)
			}
			if tt.kafkaMock != nil {
				tt.kafkaMock(mockKafka)
			}
			if tt.orderStorageMock != nil {
				tt.orderStorageMock(&s.storage)
			}

			err := s.Paid(context.Background(), tt.orderID)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())

				return
			}

			assert.NoError(t, err)
		})
	}
}
