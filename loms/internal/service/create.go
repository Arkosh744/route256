package service

import (
	"context"
	"errors"
	"time"

	"route256/libs/log"
	"route256/loms/internal/models"

	"go.uber.org/multierr"
)

func (s *service) Create(ctx context.Context, user int64, items []models.Item) (int64, error) {
	orderID, err := s.repo.CreateOrder(ctx, user)
	if err != nil {
		return 0, err
	}

	if err = s.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		var stocks []models.StockItem
		for _, item := range items {
			stocks, err = s.repo.GetStocks(ctx, item.SKU)
			if err != nil {
				return err
			}

			toReserve := uint64(item.Count)
			for _, stock := range stocks {
				if stock.Count >= toReserve {
					stockCount := stock.Count - toReserve

					if err = s.repo.CreateReservation(ctx, orderID, stock.WarehouseID, item.SKU, toReserve); err != nil {
						return err
					}

					toReserve = 0

					if stockCount == 0 {
						if err = s.repo.DeleteStock(ctx, stock.WarehouseID, item.SKU); err != nil {
							return err
						}

						break
					}

					if err = s.repo.UpdateStock(ctx, stock.WarehouseID, item.SKU, stockCount); err != nil {
						return err
					}

					break
				}

				if err = s.repo.CreateReservation(ctx, orderID, stock.WarehouseID, item.SKU, stock.Count); err != nil {
					return err
				}

				if err = s.repo.DeleteStock(ctx, stock.WarehouseID, item.SKU); err != nil {
					return err
				}

				toReserve -= stock.Count
			}

			if toReserve > 0 {
				return ErrStockInsufficient
			}

		}

		if err = s.repo.CreateOrderItems(ctx, orderID, items); err != nil {
			return err
		}

		if err = s.repo.UpdateOrderStatus(ctx, orderID, models.OrderStatusAwaitingPayment); err != nil {
			return err
		}

		s.storage.mu.Lock()
		defer s.storage.mu.Unlock()
		s.storage.storage[orderID] = &orderStatus{
			Timer: time.AfterFunc(orderTimeout, s.orderTimeoutFunc(context.Background(), orderID)),
		}

		return nil
	}); err != nil {
		if updErr := s.repo.UpdateOrderStatus(ctx, orderID, models.OrderStatusFailed); updErr != nil {
			err = multierr.Append(err, errors.New("failed to update order status to 'failed'"))
		}

		return 0, err
	}

	return orderID, nil
}

func (s *service) orderTimeoutFunc(ctx context.Context, orderID int64) func() {
	return func() {
		s.storage.mu.Lock()
		order, ok := s.storage.storage[orderID]
		s.storage.mu.Unlock()

		if ok && order.Paid {
			return
		}

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

			return
		}

		s.storage.mu.Lock()
		delete(s.storage.storage, orderID)
		s.storage.mu.Unlock()

		return
	}
}
