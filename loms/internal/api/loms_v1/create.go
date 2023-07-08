package loms_v1

import (
	"context"

	"route256/loms/internal/converter"
	desc "route256/pkg/loms_v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	var (
		userID   = req.GetUser()
		reqItems = req.GetItems()
	)

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("userID", userID)
		span.SetTag("items", reqItems)
	}

	res, err := i.lomsService.Create(ctx, userID, converter.ToItems(reqItems))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error create order: %v", err)
	}

	return &desc.CreateOrderResponse{OrderId: res}, nil
}
