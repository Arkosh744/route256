package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestsCounter       *prometheus.CounterVec
	ResponseCounter       *prometheus.CounterVec
	HistogramResponseTime *prometheus.HistogramVec
)

func Init(_ context.Context) error {
	RequestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "route256",
		Subsystem: "grpc",
		Name:      "requests_total",
	},
		[]string{"handler"},
	)

	ResponseCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "route256",
		Subsystem: "grpc",
		Name:      "responses_total",
	},
		[]string{"status", "handler"},
	)

	HistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "route256",
		Subsystem: "grpc",
		Name:      "histogram_response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	},
		[]string{"status", "handler"},
	)

	return nil
}
