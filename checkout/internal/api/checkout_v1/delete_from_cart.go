package checkout_v1

import (
	"context"

	descCheckoutV1 "route256/pkg/checkout_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteFromCart(ctx context.Context, req *descCheckoutV1.CartRequest) (*empty.Empty, error) {
	err := i.cartService.DeleteFromCart(ctx, req.GetUser(), req.GetSku(), uint16(req.GetCount()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error removing from cart: %v", err)
	}

	return &empty.Empty{}, nil
}
