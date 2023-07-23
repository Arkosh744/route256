package tg

import (
	"context"
	"route256/notifications/internal/cache"
	"route256/notifications/internal/models"
	"strconv"
)

func (b *Bot) addToCache(ctx context.Context, msg models.TelegramMessage) error {
	userID, err := b.cache.Get(ctx, cache.BuildGetUserIDByOrderKey(msg.OrderID))
	if err != nil {
		return err
	}

	var userInt int64
	if userID == "" {
		userInt, err = b.repo.GetUserIDByOrderID(ctx, msg.OrderID)
		if err != nil {
			return err
		}
	} else {
		userInt, err = strconv.ParseInt(userID, 10, 64)
		if err != nil {
			return err
		}
	}

	msg.UserID = userInt

	err = b.cache.AddToUserHistoryDay(ctx, msg.OrderMessage)
	if err != nil {
		return err
	}

	return nil
}
