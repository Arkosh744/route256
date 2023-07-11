package checkout_v1

import (
	"context"

	"route256/checkout/internal/converter"

	desc "route256/pkg/checkout_v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ListCart(ctx context.Context, req *desc.OrderIDRequest) (*desc.ListCartResponse, error) {
	userID := req.GetUser()

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("userID", userID)
	}

	res, err := i.cartService.ListCart(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error get cart: %v", err)
	}

	return converter.ListToDesc(res), nil
}
