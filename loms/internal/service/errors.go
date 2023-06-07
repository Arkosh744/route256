package service

import (
	"errors"
)

var (
	ErrStockInsufficient  = errors.New("stock insufficient")
	ErrOrderNotFound      = errors.New("order not found")
	ErrInvalidOrderStatus = errors.New("invalid order status for such operation")
)
