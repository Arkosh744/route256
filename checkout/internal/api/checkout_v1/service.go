package checkout_v1

import (
	"route256/checkout/internal/service"
	desc "route256/pkg/checkout_v1"
)

type Implementation struct {
	desc.UnimplementedCheckoutServer

	cartService service.Service
}

func NewImplementation(s service.Service) *Implementation {
	return &Implementation{
		cartService: s,
	}
}
