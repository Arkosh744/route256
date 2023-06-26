package loms_v1

import (
	"context"

	"route256/loms/internal/converter"
	desc "route256/pkg/loms_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ListOrder(ctx context.Context, req *desc.OrderIDRequest) (*desc.ListOrderResponse, error) {
	res, err := i.lomsService.Get(ctx, req.GetOrderId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error list order order: %v", err)
	}

	return converter.ToOrderDesc(res), nil
}

