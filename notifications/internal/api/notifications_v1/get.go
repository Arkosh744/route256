package notifications_v1

import (
	"context"
	"route256/notifications/internal/converter"

	desc "route256/pkg/notifications_v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ListUserHistoryDay(ctx context.Context, req *desc.ListUserHistoryDayRequest) (*desc.ListUserHistoryDayResponse, error) {
	userID := req.GetUser()

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("userID", userID)
	}

	res, err := i.service.ListUserHistoryDay(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user history order statuses: %v", err)
	}

	return &desc.ListUserHistoryDayResponse{Messages: converter.ToOrderDesc(res)}, nil
}
