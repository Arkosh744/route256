package loms_v1

import (
	"context"

	desc "route256/pkg/loms_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CancelOrder(ctx context.Context, orderID *desc.OrderIDRequest) (*empty.Empty, error) {
	if err := i.lomsService.Cancel(ctx, orderID.GetOrderId()); err != nil {
		return nil, status.Errorf(codes.Internal, "error cancel order: %v", err)
	}

	return &empty.Empty{}, nil
}
