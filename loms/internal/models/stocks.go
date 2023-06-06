package models

type StockItem struct {
	WarehouseID int64  `db:"warehouse_id" json:"warehouse_id"`
	Count       uint64 `db:"count" json:"count"`
}
