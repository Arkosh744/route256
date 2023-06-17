package loms_v1

import (
	"context"

	"route256/loms/internal/converter"
	desc "route256/pkg/loms_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	res, err := i.lomsService.Create(ctx, req.GetUser(), converter.ToItems(req.GetItems()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error create order: %v", err)
	}

	return &desc.CreateOrderResponse{OrderId: res}, nil
}
