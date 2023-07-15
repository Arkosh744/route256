package models

import "time"

type OrderMessage struct {
	UserID    int64     `db:"user_id"`
	OrderID   int64     `db:"order_id"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

type TelegramMessage struct {
	OrderMessage

	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
