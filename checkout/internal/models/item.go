package models

type ItemBase struct {
	Name  string
	Price uint32
}

type ItemInfo struct {
	ItemBase
	SKU   uint32
	Count uint32
}
