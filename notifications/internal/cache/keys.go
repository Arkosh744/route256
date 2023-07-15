package cache

import "fmt"

func BuildUserIDHistoryKey(userID int64) string {
	return fmt.Sprintf("user:%d:history:day", userID)
}

func BuildGetUserIDByOrderKey(orderID int64) string {
	return fmt.Sprintf("order:%d", orderID)
}

func BuildLatestMsgTimeKey(userID int64) string {
	return fmt.Sprintf("user:%d:latest", userID)
}