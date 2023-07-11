package interceptor

import (
	"context"
	"time"

	"route256/libs/log"
	"route256/libs/metrics"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Debug("incoming GRPC request", zap.String("method", info.FullMethod), zap.Any("request", req))
	metrics.RequestsCounter.WithLabelValues(info.FullMethod).Inc()

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			log.Debug("tracing", zap.String("traceID", sc.TraceID().String()), zap.String("spanID", sc.SpanID().String()))
		}
	}

	timeStart := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		if span := opentracing.SpanFromContext(ctx); span != nil {
			ext.Error.Set(span, true)
		}

		log.Error(ctx, "Error handling GRPC request", zap.String("method", info.FullMethod), zap.Error(err))
		metrics.ResponseCounter.WithLabelValues("error").Inc()

		elapsed := time.Since(timeStart)
		metrics.HistogramResponseTime.WithLabelValues("error").Observe(elapsed.Seconds())

		return nil, err
	}

	log.Debug("GRPC response", zap.String("method", info.FullMethod), zap.Any("response", res))
	metrics.ResponseCounter.WithLabelValues("success", info.FullMethod).Inc()

	elapsed := time.Since(timeStart)
	metrics.HistogramResponseTime.WithLabelValues("success", info.FullMethod).Observe(elapsed.Seconds())

	return res, nil
}
