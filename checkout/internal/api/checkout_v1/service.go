package checkout_v1

import (
	"route256/checkout/internal/service"
	descCheckoutV1 "route256/pkg/checkout_v1"
)

type Implementation struct {
	descCheckoutV1.UnimplementedCheckoutServer

	cartService service.Service
}

func NewImplementation(s service.Service) *Implementation {
	return &Implementation{
		cartService: s,
	}
}
