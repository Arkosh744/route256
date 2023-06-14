package service

import (
	"context"
	"errors"
	"time"

	"route256/loms/internal/models"

	"go.uber.org/multierr"
)

func (s *service) Create(ctx context.Context, user int64, items []models.Item) (int64, error) {
	orderID, err := s.repo.CreateOrder(ctx, user)
	if err != nil {
		return 0, err
	}

	if err = s.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		if txErr := s.processOrderItems(ctx, orderID, items); txErr != nil {
			return txErr
		}

		s.startPaymentTimeout(ctx, orderID)

		return nil
	}); err != nil {
		if txErr := s.repo.UpdateOrderStatus(ctx, orderID, models.OrderStatusFailed); txErr != nil {
			err = multierr.Append(err, errors.New("failed to update order status to 'failed'"))
		}

		return 0, err
	}

	return orderID, nil
}

func (s *service) processOrderItems(ctx context.Context, orderID int64, items []models.Item) error {
	if err := s.makeReservation(ctx, orderID, items); err != nil {
		return err
	}

	if err := s.repo.CreateOrderItems(ctx, orderID, items); err != nil {
		return err
	}

	if err := s.repo.UpdateOrderStatus(ctx, orderID, models.OrderStatusAwaitingPayment); err != nil {
		return err
	}

	return nil
}

func (s *service) makeReservation(ctx context.Context, orderID int64, items []models.Item) error {
	for _, item := range items {
		stocks, err := s.repo.GetStocks(ctx, item.SKU)
		if err != nil {
			return err
		}

		toReserve := uint64(item.Count)
		for _, stock := range stocks {
			if stock.Count >= toReserve {
				stockCount := stock.Count - toReserve
				if err := s.updateStockAndCreateReservation(ctx, orderID, stock.WarehouseID, item.SKU, toReserve, stockCount); err != nil {
					return err
				}

				toReserve = 0

				break
			}

			if err := s.updateStockAndCreateReservation(ctx, orderID, stock.WarehouseID, item.SKU, stock.Count, 0); err != nil {
				return err
			}

			toReserve -= stock.Count
		}

		if toReserve > 0 {
			return ErrStockInsufficient
		}
	}

	return nil
}

func (s *service) updateStockAndCreateReservation(
	ctx context.Context,
	orderID, warehouseID int64,
	sku uint32,
	toReserve, stockCount uint64,
) error {
	if err := s.repo.CreateReservation(ctx, orderID, warehouseID, sku, toReserve); err != nil {
		return err
	}

	if stockCount == 0 {
		if err := s.repo.DeleteStock(ctx, warehouseID, sku); err != nil {
			return err
		}

		return nil
	}

	if err := s.repo.UpdateStock(ctx, warehouseID, sku, stockCount); err != nil {
		return err
	}

	return nil
}

func (s *service) startPaymentTimeout(ctx context.Context, orderID int64) {
	timerCtx, cancel := context.WithCancel(ctx)

	s.storage.mu.Lock()
	defer s.storage.mu.Unlock()

	s.storage.storage[orderID] = &orderStatus{
		Timer:     time.AfterFunc(orderTimeout, s.orderTimeoutFunc(timerCtx, orderID)),
		ctxCancel: cancel,
	}
}

func (s *service) orderTimeoutFunc(ctx context.Context, orderID int64) func() {
	return func() {
		s.storage.mu.Lock()
		order, ok := s.storage.storage[orderID]
		s.storage.mu.Unlock()

		if !ok {
			return
		}

		if ok && order.Paid {
			s.storage.deleteFromStorage(orderID)

			return
		}

		if err := s.Cancel(ctx, orderID); err != nil {
			return
		}

		s.storage.deleteFromStorage(orderID)
	}
}
