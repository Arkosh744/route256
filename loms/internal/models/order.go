package models

const (
	OrderStatusNew             = "new"
	OrderStatusAwaitingPayment = "awaiting_payment"
	OrderStatusFailed          = "failed"
	OrderStatusPaid            = "paid"
	OrderStatusCanceled        = "canceled"
)

type Order struct {
	Status string `db:"status"`
	User   int64  `db:"user_id"`
	Items  []Item
}

type Item struct {
	SKU   uint32 `db:"sku"`
	Count uint32 `db:"count"`
}
