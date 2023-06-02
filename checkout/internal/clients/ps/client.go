package ps

import (
	"context"

	"route256/checkout/internal/config"
	"route256/checkout/internal/converter"
	"route256/checkout/internal/models"
	productV1 "route256/pkg/product_v1"
)

var _ Client = (*client)(nil)

type Client interface {
	GetProduct(ctx context.Context, sku uint32) (*models.ItemBase, error)
}

type client struct {
	psClient productV1.ProductServiceClient
}

func New(ps productV1.ProductServiceClient) *client {
	return &client{
		psClient: ps,
	}
}

func (c *client) GetProduct(ctx context.Context, sku uint32) (*models.ItemBase, error) {
	res, err := c.psClient.GetProduct(ctx, &productV1.GetProductRequest{
		Token: config.AppConfig.Token,
		Sku:   sku,
	})
	if err != nil {
		return nil, err
	}

	return converter.DescToItemBase(res), nil
}
