package models

const (
	OrderStatysUnknown         = "unknown"
	OrderStatusNew             = "new"
	OrderStatusAwaitingPayment = "awaiting_payment"
	OrderStatusFailed          = "failed"
	OrderStatusPaid            = "paid"
	OrderStatusCanceled        = "canceled"
)

type Order struct {
	Status string
	User   int64
	Items  []Item
}

type Item struct {
	SKU   uint32
	Count uint32
}
