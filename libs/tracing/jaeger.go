package tracing

import (
	"github.com/uber/jaeger-client-go/config"
)

func Init(host, serviceName string) error{
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			// LocalAgentHostPort: "jaeger:6831",
			LocalAgentHostPort: host,
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		return err
	}

	return nil
}
