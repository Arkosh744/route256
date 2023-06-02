package checkout_v1

import (
	"context"

	desc "route256/pkg/checkout_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Purchase(ctx context.Context, req *desc.OrderIDRequest) (*desc.OrderIDResponse, error) {
	orderID, err := i.cartService.Purchase(ctx, req.GetUser())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error purchase: %v", err)
	}

	return &desc.OrderIDResponse{OrderId: orderID}, nil
}
