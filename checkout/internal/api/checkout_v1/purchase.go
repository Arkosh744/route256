package checkout_v1

import (
	"context"

	descCheckoutV1 "route256/pkg/checkout_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Purchase(ctx context.Context, req *descCheckoutV1.OrderIDRequest) (*descCheckoutV1.OrderIDResponse, error) {
	_, err := i.cartService.Purchase(ctx, req.GetUser())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error purchase: %v", err)
	}

	return &descCheckoutV1.OrderIDResponse{}, nil
}
