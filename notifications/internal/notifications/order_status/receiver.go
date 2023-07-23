package order_status

import (
	"context"
	"fmt"
	"log"
	"route256/notifications/internal/models"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

type OrderStatusReceiver interface {
	Subscribe(topic string) error
}

type MessageSender interface {
	SendMessage(ctx context.Context, tgMessage models.TelegramMessage) error
}

type receiver struct {
	consumer sarama.Consumer
	bot      MessageSender
	chatID   int64
}

func NewReceiver(consumer sarama.Consumer, bot MessageSender, chatID int64) *receiver {
	return &receiver{
		consumer: consumer,
		bot:      bot,
		chatID:   chatID,
	}
}

func (r *receiver) Subscribe(ctx context.Context, topic string) error {
	partitionList, err := r.consumer.Partitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitionList {
		pc, err := r.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}

		go r.processMessages(ctx, pc, topic)
	}

	return nil
}

func (r *receiver) processMessages(ctx context.Context, pc sarama.PartitionConsumer, topic string) {
	for message := range pc.Messages() {
		orderID := string(message.Key)
		status := string(message.Value)

		messageText := fmt.Sprintf("New message:\norderID: `%s`\nstatus: `%s`\ntopic: %s    partion: %d    offset: %d",
			orderID,
			status,
			topic,
			message.Partition,
			message.Offset)

		log.Println(messageText)

		orderInt, err := strconv.ParseInt(orderID, 10, 64)
		if err != nil {
			log.Printf("failed to convert orderID to int: %v", err)
			continue
		}

		tgMessage := models.TelegramMessage{
			OrderMessage: models.OrderMessage{
				OrderID: orderInt,
				Status:  status,
				CreatedAt: time.Now(),
			},
			ChatID: r.chatID,
			Text:   messageText,
		}

		if err := r.bot.SendMessage(ctx, tgMessage); err != nil {
			log.Printf("failed to send message: %v", err)
		}
	}
}
