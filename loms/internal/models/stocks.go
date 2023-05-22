package models

type StockItem struct {
	WarehouseID int64  `json:"warehouse_id"`
	Count       uint64 `json:"count"`
}
