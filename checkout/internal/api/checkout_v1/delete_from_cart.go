package checkout_v1

import (
	"context"

	desc "route256/pkg/checkout_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteFromCart(ctx context.Context, req *desc.CartRequest) (*empty.Empty, error) {
	if err := i.cartService.DeleteFromCart(ctx, req.GetUser(), req.GetSku(), uint16(req.GetCount())); err != nil {
		return nil, status.Errorf(codes.Internal, "error removing from cart: %v", err)
	}

	return &empty.Empty{}, nil
}
