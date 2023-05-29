package get

import (
	"context"
	"errors"

	"route256/loms/internal/log"
	"route256/loms/internal/models"
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
	User int64 `json:"user"`
}

type Response struct {
	*models.Order
}

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrUserNotFound
	}

	return nil
}

var ErrUserNotFound = errors.New("user not found")

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Infof("%+v", req)

	res, err := h.service.Get(ctx, req.User)
	if err != nil {
		return Response{}, err
	}

	return Response{res}, err
}
