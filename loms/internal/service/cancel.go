package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"route256/libs/log"
	"route256/loms/internal/models"
)

func (s *service) Cancel(ctx context.Context, orderID int64) error {
	if err := s.cancelOrderAndRestock(ctx, orderID); err != nil {
		if errors.Is(err, ErrOrderNotFound) || errors.Is(err, ErrInvalidOrderStatus) {
			return err
		}

		log.Errorf("failed to cancel order %d: %v", orderID, err)

		return err
	}

	storageOrder, ok := s.storage.storage[orderID]
	if ok {
		storageOrder.ctxCancel()
		s.storage.deleteFromStorage(orderID)
	}

	if err := s.kafka.SendOrderStatus(orderID, models.OrderStatusCanceled); err != nil {
		log.Errorf("failed to send order status: %v", err)

		return err
	}

	return nil
}

func (s *service) cancelOrderAndRestock(ctx context.Context, orderID int64) error {
	if err := s.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		order, err := s.repo.GetOrder(ctx, orderID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ErrOrderNotFound
			}

			return err
		}

		if order.Status != models.OrderStatusNew && order.Status != models.OrderStatusAwaitingPayment {
			return ErrInvalidOrderStatus
		}

		if err = s.repo.UpdateOrderStatus(ctx, orderID, models.OrderStatusCanceled); err != nil {
			return fmt.Errorf("failed to update order %d status to 'canceled': %w", orderID, err)
		}

		reserv, err := s.repo.GetReservations(ctx, orderID)
		if err != nil {
			return fmt.Errorf("failed to get reservations for order %d: %w", orderID, err)
		}

		for i := range reserv {
			if err = s.repo.InsertStock(ctx, reserv[i]); err != nil {
				return fmt.Errorf("failed to insert stock back from order %d: %w", orderID, err)
			}
		}

		if err = s.repo.DeleteReservation(ctx, orderID); err != nil {
			return fmt.Errorf("failed to delete reservation for order %d: %v", orderID, err)
		}

		return nil
	}); err != nil {

		return err
	}

	return nil
}
