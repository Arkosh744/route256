package handlers

import (
	"net/http"
	"route256/checkout/internal/handlers/add"
	"route256/checkout/internal/handlers/del"
	"route256/checkout/internal/handlers/get"
	"route256/checkout/internal/handlers/purchase"
	"route256/checkout/internal/service"
	"route256/libs/wrappers"
)

func InitRouter(service service.Service) *http.ServeMux {
	mux := http.NewServeMux()

	addToCart := add.NewHandler(service).Handle
	mux.Handle("/addToCart/", wrappers.New(addToCart))

	deleteFromCart := del.NewHandler(service).Handle
	mux.Handle("/deleteFromCart/", wrappers.New(deleteFromCart))

	listCart := get.NewHandler(service).Handle
	mux.Handle("/listCart/", wrappers.New(listCart))

	buy := purchase.NewHandler(service).Handle
	mux.Handle("/purchase/", wrappers.New(buy))

	return mux
}
