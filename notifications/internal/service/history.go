package service

import (
	"context"
	"route256/notifications/internal/models"
	"time"
)

func (s *Service) ListUserHistoryDay(ctx context.Context, userID int64) ([]models.OrderMessage, error) {
	msgs, err := s.cache.GetUserHistoryDay(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(msgs) > 0 {
		var lastMsgTime time.Time
		lastMsgTime, err = s.cache.GetLatestMessageTime(ctx, userID)
		if err != nil {
			return nil, err
		}

		var msgsNew []models.OrderMessage
		msgsNew, err = s.repo.ListUserHistoryDay(ctx, userID, &lastMsgTime)
		if err != nil {
			return nil, err
		}

		if len(msgsNew) > 0 {
			msgs = append(msgs, msgsNew...)
			for _, msg := range msgsNew {
				msg.UserID = userID
				if err = s.cache.AddToUserHistoryDay(ctx, msg); err != nil {
					return nil, err
				}
			}
		}

		return msgs, nil
	}

	msgs, err = s.repo.ListUserHistoryDay(ctx, userID, nil)
	if err != nil {
		return nil, err
	}

	for _, msg := range msgs {
		msg.UserID = userID
		if err = s.cache.AddToUserHistoryDay(ctx, msg); err != nil {
			return nil, err
		}
	}

	return msgs, nil
}
