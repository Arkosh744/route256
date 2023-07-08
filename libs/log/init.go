package log

import (
	"context"
	l "log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

const (
	productionPreset = "prod"
	devPreset        = "dev"
)

func InitLogger(_ context.Context, preset string) error {
	zapLog, err := selectLogger(preset)
	if err != nil {
		return err
	}

	log = zapLog

	return nil
}

func selectLogger(preset string) (*zap.Logger, error) {
	cfg := getConfig(preset)

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func getConfig(preset string) zap.Config {
	var cfg zap.Config

	switch preset {
	case devPreset:
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    customLevelEncoder,
			EncodeTime:     customTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {},
		}

		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig = encoderConfig

		return cfg
	case productionPreset:
		cfg = zap.NewProductionConfig()
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

		return cfg
	default:
		l.Println("unknown logger preset, using development preset")

		return zap.NewDevelopmentConfig()
	}
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05") + "  |")
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString() + "  |")
}
