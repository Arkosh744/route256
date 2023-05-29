package loms

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"route256/checkout/internal/models"
	"route256/libs/wrappers"
)

type StocksRequest struct {
	SKU uint32 `json:"sku"`
}

type StocksResponse struct {
	Stocks []struct {
		WarehouseID int64  `json:"warehouse_id"`
		Count       uint64 `json:"count"`
	} `json:"stocks"`
}

func (c *client) Stocks(ctx context.Context, sku uint32) ([]*models.Stock, error) {
	req := StocksRequest{SKU: sku}

	reqPath, err := url.JoinPath(c.host, stocksPath)
	if err != nil {
		return nil, err
	}

	res, err := wrappers.Do[StocksRequest, StocksResponse](ctx, &req, http.MethodPost, reqPath)
	if err != nil {
		return nil, fmt.Errorf(`do request "%s": %w`, reqPath, err)
	}

	result := make([]*models.Stock, 0, len(res.Stocks))
	for _, v := range res.Stocks {
		result = append(result, &models.Stock{
			WarehouseID: v.WarehouseID,
			Count:       v.Count,
		})
	}

	return result, nil
}
