package loms_v1

import (
	"route256/loms/internal/service"
	desc "route256/pkg/loms_v1"
)

type Implementation struct {
	desc.UnimplementedLomsServer

	lomsService service.Service
}

func NewImplementation(s service.Service) *Implementation {
	return &Implementation{
		lomsService: s,
	}
}
