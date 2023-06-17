package models

type ItemInfo struct {
	Name  string
	Price uint32
}

type ItemData struct {
	SKU   uint32 `db:"sku"`
	Count uint16 `db:"count"`
}

type Item struct {
	ItemInfo
	ItemData
}

type CartInfo struct {
	Items      []Item
	TotalPrice uint32
}
