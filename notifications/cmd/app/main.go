package main

import (
	"context"
	l "log"
	"os"
	"os/signal"
	"route256/notifications/internal/app"
	"syscall"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		l.Fatalf("failed to initialize app: %v", err)
	}

	if err = a.Run(ctx); err != nil {
		l.Fatalf("failed to run app: %v", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
}
