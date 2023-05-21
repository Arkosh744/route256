package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	l "log"
	"route256/checkout/internal/config"
	"time"
)

var log *zap.SugaredLogger

const (
	productionPreset = "prod"
	devPreset        = "dev"
)

func InitLogger(_ context.Context) error {
	zapLog, err := selectLogger()
	if err != nil {
		return err
	}

	log = zapLog.Sugar()

	return nil
}

func selectLogger() (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

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
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var cfg zap.Config
	switch config.AppConfig.Log.Preset {
	case productionPreset:
		cfg = zap.NewProductionConfig()
	case devPreset:
		cfg = zap.NewDevelopmentConfig()
	default:
		l.Println("unknown logger preset, using development preset")
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.EncoderConfig = encoderConfig
	logger, err = cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05") + "  |")
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString() + "  |")
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}
