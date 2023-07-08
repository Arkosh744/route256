package log

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
)

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	span := opentracing.SpanFromContext(ctx)

	if span != nil {
		if spancontext, ok := span.Context().(jaeger.SpanContext); ok {
			fields = append(
				fields,
				zap.String("trace", spancontext.TraceID().String()),
				zap.String("span", spancontext.SpanID().String()),
			)
		}
	}

	log.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}
