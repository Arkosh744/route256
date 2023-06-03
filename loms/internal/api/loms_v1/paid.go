package loms_v1

import (
	"context"

	desc "route256/pkg/loms_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) OrderPaid(ctx context.Context, orderID *desc.OrderIDRequest) (*empty.Empty, error) {
	if err := i.lomsService.Paid(ctx, orderID.GetOrderId()); err != nil {
		return nil, status.Errorf(codes.Internal, "error set order paid order: %v", err)
	}

	return &empty.Empty{}, nil
}
