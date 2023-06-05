package models

type ItemBase struct {
	Name  string
	Price uint32
}

type ItemData struct {
	SKU   uint32 `db:"sku"`
	Count uint32 `db:"count"`
}

type ItemInfo struct {
	ItemBase
	ItemData
}

type CartInfo struct {
	Items      []ItemInfo
	TotalPrice uint32
}
