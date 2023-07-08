package tracing

import (
	"route256/libs/log"

	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

func Init(host, serviceName string) error {
	log.Info("Initializing tracing", zap.String("host", host), zap.String("service", serviceName))

	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort: host,
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		return err
	}

	return nil
}
