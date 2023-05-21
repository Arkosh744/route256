package main

import (
	"context"
	"log"
	"route256/checkout/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	if err = a.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
