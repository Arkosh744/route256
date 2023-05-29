package create

import (
	"context"

	"route256/loms/internal/log"
	"route256/loms/internal/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

type Request struct {
	User  int64  `json:"user"`
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	OrderID int64 `json:"order_id"`
}

func (r Request) Validate() error {
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Infof("%+v", req)

	orderID, err := h.service.Create(ctx, req.User, req.SKU, req.Count)
	if err != nil {
		return Response{}, err
	}

	return Response{orderID}, nil
}
