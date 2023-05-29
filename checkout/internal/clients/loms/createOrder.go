package loms

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"route256/checkout/internal/models"
	"route256/libs/wrappers"
)

type CreateOrderRequest struct {
	User  int64              `json:"user"`
	Items []*models.ItemData `json:"items"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"order_id"`
}

func (c *client) CreateOrder(ctx context.Context, user int64, items []*models.ItemData) (int64, error) {
	req := CreateOrderRequest{User: user, Items: items}

	reqPath, err := url.JoinPath(c.host, createOrderPath)
	if err != nil {
		return 0, err
	}

	res, err := wrappers.Do[CreateOrderRequest, CreateOrderResponse](ctx, &req, http.MethodPost, reqPath)
	if err != nil {
		return 0, fmt.Errorf(`do request "%s": %w`, reqPath, err)
	}

	return res.OrderID, nil
}
