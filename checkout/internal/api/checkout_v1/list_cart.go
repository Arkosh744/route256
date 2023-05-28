package checkout_v1

import (
	"context"

	descCheckoutV1 "route256/pkg/checkout_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ListCart(ctx context.Context, req *descCheckoutV1.OrderIDRequest) (*descCheckoutV1.ListCartResponse, error) {
	_, err := i.cartService.ListCart(ctx, req.GetUser())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error get cart: %v", err)
	}

	return &descCheckoutV1.ListCartResponse{}, nil
}
