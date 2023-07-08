package loms_v1

import (
	"context"
	"errors"

	"route256/loms/internal/service"
	desc "route256/pkg/loms_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CancelOrder(ctx context.Context, orderID *desc.OrderIDRequest) (*empty.Empty, error) {
	id := orderID.GetOrderId()

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("orderID", id)
	}

	if err := i.lomsService.Cancel(ctx, id); err != nil {
		if errors.Is(err, service.ErrOrderNotFound) {
			return nil, status.Errorf(codes.NotFound, "order not found: %v", err)
		}

		if errors.Is(err, service.ErrInvalidOrderStatus) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid order status: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "error cancel order: %v", err)
	}

	return &empty.Empty{}, nil
}
