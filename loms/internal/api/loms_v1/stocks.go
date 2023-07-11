package loms_v1

import (
	"context"

	"route256/loms/internal/converter"
	desc "route256/pkg/loms_v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Stocks(ctx context.Context, req *desc.StocksRequest) (*desc.StocksResponse, error) {
	sku := req.GetSku()

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("SKU", sku)
	}

	res, err := i.lomsService.Stocks(ctx, sku)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error get Stocks: %v", err)
	}

	return &desc.StocksResponse{Stocks: converter.ToStocksDesc(res)}, nil
}
