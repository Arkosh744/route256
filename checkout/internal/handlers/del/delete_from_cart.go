package del

import (
	"context"
	"errors"
	"route256/checkout/internal/log"
	"route256/checkout/internal/service"
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

type Response struct{}

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrUserNotFound
	}

	return nil
}

var ErrUserNotFound = errors.New("user not found")

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Infof("%+v", req)

	err := h.service.DeleteFromCart(ctx, req.User, req.SKU, req.Count)

	return Response{}, err
}