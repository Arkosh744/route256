package loms_v1

import (
	"context"

	desc "route256/pkg/loms_v1"

	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) OrderPayed(ctx context.Context, orderID *desc.OrderIDRequest) (*empty.Empty, error) {
	err := i.lomsService.Paid(ctx, orderID.GetOrderId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
