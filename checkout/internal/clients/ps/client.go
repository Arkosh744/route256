package ps

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/converter"
	"route256/checkout/internal/models"
	"route256/libs/log"
	"route256/libs/rate_limiter"
	productV1 "route256/pkg/product_v1"
	"time"
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

func (c *client) GetProduct(ctx context.Context, sku uint32) (*models.ItemInfo, error) {
	log.Infof("get product from ps: sku %d", sku)

	for {
		if c.rl.Allow() {
			break
		}

		// Because we have a sliding window, we will wait for a half of the period and retry
		time.Sleep(config.AppConfig.ReqLimitPeriod / 2)
	}

	res, err := c.psClient.GetProduct(ctx, &productV1.GetProductRequest{
		Token: config.AppConfig.Token,
		Sku:   sku,
	})
	if err != nil {
		return nil, err
	}

	return converter.DescToItemBase(res), nil
}
