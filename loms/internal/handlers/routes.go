package handlers

import (
	"net/http"
	"route256/libs/wrappers"
	"route256/loms/internal/handlers/create"
	"route256/loms/internal/handlers/get"
	"route256/loms/internal/handlers/paid"
	"route256/loms/internal/handlers/stocks"
	"route256/loms/internal/service"
)

func InitRouter(service service.Service) *http.ServeMux {
	mux := http.NewServeMux()

	createOrder := create.NewHandler(service).Handle
	mux.Handle("/createOrder", wrappers.New(createOrder))

	listOrder := get.NewHandler(service).Handle
	mux.Handle("/listOrder", wrappers.New(listOrder))

	orderPaid := paid.NewHandler(service).Handle
	mux.Handle("/orderPaid", wrappers.New(orderPaid))

	cancelOrder := paid.NewHandler(service).Handle
	mux.Handle("/cancelOrder", wrappers.New(cancelOrder))

	stock := stocks.NewHandler(service).Handle
	mux.Handle("/stocks", wrappers.New(stock))

	return mux
}
