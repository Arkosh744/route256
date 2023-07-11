package service

import (
	"context"
	"testing"

	"route256/libs/client/pg"
	"route256/loms/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_service_Create(t *testing.T) {
	ctx := context.Background()
	userID := int64(1)
	items := []models.Item{
		{
			SKU:   123,
			Count: 10,
		},
	}

	tests := []struct {
        name             string
        repoMock         func(m *MockRepository)
        txManagerMock    func(m *pg.MockTxManager)
        kafkaMock        func(m *MockOrderStatusSender)
        wantErr          error
    }{
        {
            name: "success",
            repoMock: func(m *MockRepository) {
                m.EXPECT().CreateOrder(ctx, userID).Return(int64(1), nil).Times(1)
                m.EXPECT().GetStocks(ctx, items[0].SKU).Return(
                    []models.StockItem{{WarehouseID: 1, Count: uint64(items[0].Count)}}, nil).Times(1)
                m.EXPECT().CreateReservation(ctx, int64(1), int64(1), items[0].SKU, uint64(items[0].Count)).Return(nil).Times(1)
                m.EXPECT().DeleteStock(ctx, int64(1), items[0].SKU).Return(nil).Times(1)
                m.EXPECT().CreateOrderItems(ctx, int64(1), items).Return(nil).Times(1)
                m.EXPECT().UpdateOrderStatus(ctx, int64(1), models.OrderStatusAwaitingPayment).Return(nil).Times(1)
            },
            txManagerMock: func(m *pg.MockTxManager) {
                m.EXPECT().RunRepeatableRead(ctx, gomock.Any()).DoAndReturn(
                    func(ctx context.Context, f func(context.Context) error) error {
                        return f(ctx)
                    }).Times(1)
            },
            kafkaMock: func(m *MockOrderStatusSender) {
                m.EXPECT().SendOrderStatus(int64(1), models.OrderStatusNew).Return(nil).Times(1)
                m.EXPECT().SendOrderStatus(int64(1), models.OrderStatusAwaitingPayment).Return(nil).Times(1)
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

			_, err := s.Create(context.Background(), userID, items)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())

				return
			}

			assert.NoError(t, err)
		})
	}
}
