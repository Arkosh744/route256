package models

type ItemBase struct {
	Name  string
	Price uint32
}

type ItemData struct {
	SKU   uint32
	Count uint32
}

type ItemInfo struct {
	ItemBase
	ItemData
}

type CartInfo struct {
	Items      []ItemInfo `json:"items"`
	TotalPrice uint32     `json:"total_price"`
}
