package ps

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"route256/checkout/internal/config"
	"route256/checkout/internal/models"
	"route256/libs/wrappers"
)

const (
	getProductPath = "get_product"
)

type Client interface {
	GetProduct(ctx context.Context, sku uint32) (*models.ItemBase, error)
}

type client struct {
	pathStock string
}

type ItemsRequest struct {
	Token string `json:"token"`
	SKU   uint32 `json:"sku"`
}

func New(clientURL string) *client {
	stockURL, _ := url.JoinPath(clientURL, getProductPath)

	return &client{pathStock: stockURL}
}

func (c *client) GetProduct(ctx context.Context, sku uint32) (*models.ItemBase, error) {
	req := ItemsRequest{Token: config.AppConfig.Token, SKU: sku}

	res, err := wrappers.Do[ItemsRequest, models.ItemBase](ctx, &req, http.MethodPost, c.pathStock)
	if err != nil {
		return nil, fmt.Errorf(`do request "%s": %w`, c.pathStock, err)
	}

	return res, nil
}
