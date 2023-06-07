package service

import (
	"context"

	"route256/libs/log"
	"route256/loms/internal/models"
)

func (s *service) Cancel(ctx context.Context, orderID int64) error {
	if err := s.cancelOrderAndRestock(ctx, orderID); err != nil {
		return err
	}

	storageOrder, ok := s.storage.storage[orderID]
	if ok {
		storageOrder.ctxCancel()
		s.storage.deleteFromStorage(orderID)
	}

	return nil
}

func (s *service) cancelOrderAndRestock(ctx context.Context, orderID int64) error {
	if err := s.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		if err := s.repo.UpdateOrderStatus(ctx, orderID, models.OrderStatusCanceled); err != nil {
			log.Errorf("failed to update order %d status to 'canceled': %v", orderID, err)

			return err
		}

		reserv, err := s.repo.GetReservations(ctx, orderID)
		if err != nil {
			log.Errorf("failed to get reservations for order %d: %v", orderID, err)

			return err
		}

		for i := range reserv {
			if err = s.repo.InsertStock(ctx, reserv[i]); err != nil {
				log.Errorf("failed to insert stock back from order %d: %v", orderID, err)

				return err
			}
		}

		if err = s.repo.DeleteReservation(ctx, orderID); err != nil {
			log.Errorf("failed to delete reservation for order %d: %v", orderID, err)

			return err
		}

		return nil
	}); err != nil {
		log.Errorf("failed to cancel order %d: %v", orderID, err)

		return err
	}

	return nil
}
