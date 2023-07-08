package loms_v1

import (
	"context"

	desc "route256/pkg/loms_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) OrderPaid(ctx context.Context, orderID *desc.OrderIDRequest) (*empty.Empty, error) {
	id := orderID.GetOrderId()

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("orderID", id)
	}

	if err := i.lomsService.Paid(ctx, id); err != nil {
		return nil, status.Errorf(codes.Internal, "error set order paid order: %v", err)
	}

	return &empty.Empty{}, nil
}
