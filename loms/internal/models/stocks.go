package models

type StockItem struct {
	WarehouseID int64  `db:"warehouse_id" json:"warehouse_id"`
	Count       uint64 `db:"count" json:"count"`
}

type ReservationItem struct {
	StockItem
	SKU uint32 `db:"sku" json:"sku"`
}
