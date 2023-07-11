package loms_v1

import (
	"context"

	"route256/loms/internal/converter"
	desc "route256/pkg/loms_v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ListOrder(ctx context.Context, req *desc.OrderIDRequest) (*desc.ListOrderResponse, error) {
	orderID := req.GetOrderId()

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("orderID", orderID)
	}

	res, err := i.lomsService.Get(ctx, orderID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error list order order: %v", err)
	}

	return converter.ToOrderDesc(res), nil
}
