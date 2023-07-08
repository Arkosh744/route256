//nolint:dupl // Similar to other methods
package checkout_v1

import (
	"context"

	desc "route256/pkg/checkout_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) AddToCart(ctx context.Context, req *desc.CartRequest) (*empty.Empty, error) {
	var (
		userID = req.GetUser()
		sku    = req.GetSku()
		count  = uint16(req.GetCount())
	)

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("userID", userID)
		span.SetTag("SKU", sku)
		span.SetTag("count", count)
	}

	if err := i.cartService.AddToCart(ctx, userID, sku, count); err != nil {
		return nil, status.Errorf(codes.Internal, "error adding to cart: %v", err)
	}

	return &empty.Empty{}, nil
}
