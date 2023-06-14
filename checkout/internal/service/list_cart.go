package service

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/models"
	wp "route256/libs/worker_pool"
)

func (s *cartService) ListCart(ctx context.Context, user int64) (*models.CartInfo, error) {
	userItems, err := s.repo.GetUserCart(ctx, user)
	if err != nil {
		return nil, err
	}

	items := make([]models.Item, 0, len(userItems))

	ctx, cancel := context.WithCancel(ctx)

	pool := wp.NewPool[models.ItemData, models.Item](ctx, config.AppConfig.Workers)

	pool.SubmitMany(func(ctx context.Context, item models.ItemData) (models.Item, error) {
		res, err := s.psClient.GetProduct(ctx, item.SKU)
		if err != nil {
			// if we get error from PS, we cancel all other requests
			cancel()
			return models.Item{}, err
		}

		resItem := models.Item{
			ItemInfo: models.ItemInfo{
				Name:  res.Name,
				Price: res.Price,
			},
			ItemData: models.ItemData{
				SKU:   item.SKU,
				Count: item.Count,
			},
		}

		return resItem, nil
	}, userItems)

	pool.Wait()

	results := pool.GetResult()

	var totalPrice uint32
	for i := range results {
		if results[i].Err != nil {
			// if we get any error from PS, we return error
			return nil, results[i].Err
		}

		items = append(items, results[i].Value)

		totalPrice += results[i].Value.Price * uint32(results[i].Value.Count)
	}

	return &models.CartInfo{
		Items:      items,
		TotalPrice: totalPrice,
	}, nil
}
