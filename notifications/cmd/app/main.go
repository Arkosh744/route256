package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"route256/notifications/internal/config"
	"route256/notifications/internal/kafka"
	orderStatus "route256/notifications/internal/notifications/order_status"
	"route256/notifications/internal/tg"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("Unable to init config: %v", err)
	}

	consumer, err := kafka.NewConsumer(config.AppConfig.Kafka.Brokers)
	if err != nil {
		log.Fatalf("Unable to create kafka consumer: %v", err)
	}

	bot, err := tg.NewBot(config.AppConfig.Tg.Token)
	if err != nil {
		log.Fatalf("Unable to create telegram bot: %v", err)
	}

	receiver := orderStatus.NewReceiver(consumer, bot, config.AppConfig.Tg.ChatID)
	if err = receiver.Subscribe(config.AppConfig.Kafka.Topic); err != nil {
		log.Fatalf("Unable to subscribe to kafka topic: %v", err)
	}

	// Graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	if err = consumer.Close(); err != nil {
		log.Fatalf("Failed to close consumer: %v", err)
	}
}
