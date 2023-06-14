package ps

import (
	"context"

	"route256/checkout/internal/config"
	"route256/checkout/internal/converter"
	"route256/checkout/internal/models"
	"route256/libs/log"
	"route256/libs/rate_limiter"
	wp "route256/libs/worker_pool"
	productV1 "route256/pkg/product_v1"
)

type client struct {
	psClient productV1.ProductServiceClient
	rl       *rate_limiter.SlidingWindow
}

func New(ps productV1.ProductServiceClient, rl *rate_limiter.SlidingWindow) *client {
	return &client{
		psClient: ps,
		rl:       rl,
	}
}

func (c *client) getProduct(ctx context.Context, sku uint32) (*models.ItemInfo, error) {
	log.Infof("get product from ps: sku %d", sku)

	// waiting for allow from rate limiter
	c.rl.Wait()

	res, err := c.psClient.GetProduct(ctx, &productV1.GetProductRequest{
		Token: config.AppConfig.Token,
		Sku:   sku,
	})
	if err != nil {
		return nil, err
	}

	return converter.DescToItemBase(res), nil
}

func (c *client) GetProducts(ctx context.Context, userItems []models.ItemData) []wp.Result[models.Item] {
	ctx, cancel := context.WithCancel(ctx)

	pool := wp.NewPool[models.ItemData, models.Item](ctx, config.AppConfig.Workers)

	pool.SendMany(func(ctx context.Context, item models.ItemData) (models.Item, error) {
		res, err := c.getProduct(ctx, item.SKU)
		if err != nil {
			// if we get error from PS, we cancel all other requests too
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

	return pool.GetResult()
}
