package checkout_v1

import (
	"context"

	"route256/checkout/internal/converter"

	desc "route256/pkg/checkout_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ListCart(ctx context.Context, req *desc.OrderIDRequest) (*desc.ListCartResponse, error) {
	res, err := i.cartService.ListCart(ctx, req.GetUser())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error get cart: %v", err)
	}

	return converter.ListToDesc(res), nil
}
