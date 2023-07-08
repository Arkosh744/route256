package main

import (
	"context"
	"fmt"
	l "log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"route256/libs/log"
	"route256/libs/metrics"
	"route256/notifications/internal/config"
	"route256/notifications/internal/kafka"
	orderStatus "route256/notifications/internal/notifications/order_status"
	"route256/notifications/internal/tg"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	if err := config.Init(); err != nil {
		l.Fatalf("Unable to init config: %v", err)
	}

	if err := log.InitLogger(ctx, config.AppConfig.Log.Preset); err != nil {
		l.Fatalf("Unable to init logger: %v", err)
	}

	consumer, err := kafka.NewConsumer(config.AppConfig.Kafka.Brokers)
	if err != nil {
		log.Fatal("Unable to create kafka consumer", zap.Error(err))
	}

	bot, err := tg.NewBot(config.AppConfig.Tg.Token)
	if err != nil {
		log.Fatal("Unable to create telegram bot", zap.Error(err))
	}

	receiver := orderStatus.NewReceiver(consumer, bot, config.AppConfig.Tg.ChatID)
	if err = receiver.Subscribe(config.AppConfig.Kafka.Topic); err != nil {
		log.Fatal("Unable to subscribe to kafka topic", zap.Error(err))
	}

	wg := sync.WaitGroup{}

	wg.Add(1)

	if err = metrics.Init(ctx); err != nil {
		log.Fatal("Unable to init metrics", zap.Error(err))
	}

	go func() {
		defer wg.Done()

		if err = RunPrometheusServer(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Unable to run prometheus server", zap.Error(err))
		}
	}()

	// Graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	if err = consumer.Close(); err != nil {
		log.Fatal("Failed to close consumer", zap.Error(err))
	}
}

func RunPrometheusServer() error {
	metricsServer := &http.Server{
		Addr: config.AppConfig.GetMetricsAddr(),
	}

	http.Handle("/metrics", promhttp.Handler())

	log.Info(fmt.Sprintf("Metrics server is running on %s", config.AppConfig.GetMetricsAddr()))

	err := metricsServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
