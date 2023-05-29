package stocks

import (
	"context"

	"route256/loms/internal/log"
	"route256/loms/internal/models"
	"route256/loms/internal/service"
)

type Handler struct {
	service service.Service
}

type Response struct {
	Stocks []models.StockItem `json:"stocks"`
}

func NewHandler(serv service.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

type Request struct {
	SKU uint32 `json:"sku"`
}

func (r Request) Validate() error {
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Infof("%+v", req)

	return Response{
		Stocks: []models.StockItem{
			{WarehouseID: 1, Count: 200},
			{WarehouseID: 2131, Count: 3},
		},
	}, nil
}
