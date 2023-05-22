package main

import (
	"context"
	"log"
	"route256/loms/internal/app"
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

	//hand := &stocks.Handler{}
	//http.Handle("/stocks", wrappers.New(hand.Handle))
	//err := http.ListenAndServe(port, nil)
	//if err != nil {
	//	log.Fatalln("ERR: ", err)
	//}
}
