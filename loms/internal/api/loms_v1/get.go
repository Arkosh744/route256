package loms_v1

import (
	"context"

	"route256/loms/internal/converter"
	desc "route256/pkg/loms_v1"
)

func (i *Implementation) ListOrder(ctx context.Context, req *desc.OrderIDRequest) (*desc.ListOrderResponse, error) {
	res, err := i.lomsService.Get(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}

	return converter.ToOrderDesc(res), nil
}
