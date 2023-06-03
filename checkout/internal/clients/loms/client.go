package loms

import (
	lomsV1 "route256/pkg/loms_v1"
)

type client struct {
	lomsClient lomsV1.LomsClient
}

func New(loms lomsV1.LomsClient) *client {
	return &client{
		lomsClient: loms,
	}
}
