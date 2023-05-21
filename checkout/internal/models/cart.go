package models

type CartInfo struct {
	Items      []ItemInfo `json:"items"`
	TotalPrice uint32     `json:"total_price"`
}
