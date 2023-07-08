package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"route256/libs/log"
	"route256/loms/internal/models"

	"go.uber.org/zap"
)

func (s *service) Paid(ctx context.Context, orderID int64) error {
	if err := s.payOrder(ctx, orderID); err != nil {
		return err
	}

	storageOrder, ok := s.storage.storage[orderID]
	if ok {
		storageOrder.ctxCancel()
		s.storage.deleteFromStorage(orderID)
	}

	if err := s.kafka.SendOrderStatus(orderID, models.OrderStatusPaid); err != nil {
		log.Error(ctx, "failed to send order status", zap.Error(err))

		return err
	}

	return nil
}

func (s *service) payOrder(ctx context.Context, orderID int64) error {
	if err := s.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		order, err := s.repo.GetOrder(ctx, orderID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ErrOrderNotFound
			}

			return err
		}

		if order.Status != models.OrderStatusAwaitingPayment {
			return ErrInvalidOrderStatus
		}

		if err := s.repo.UpdateOrderStatus(ctx, orderID, models.OrderStatusPaid); err != nil {
			log.Error(ctx, fmt.Sprintf("failed to update order %d status to 'canceled'", orderID), zap.Error(err))

			return err
		}

		if err := s.repo.DeleteReservation(ctx, orderID); err != nil {
			log.Error(ctx, fmt.Sprintf("failed to delete reservation for order %d", orderID), zap.Error(err))

			return err
		}

		return nil
	}); err != nil {
		log.Error(ctx, fmt.Sprintf("failed to cancel order %d", orderID), zap.Error(err))

		return err
	}

	return nil
}
