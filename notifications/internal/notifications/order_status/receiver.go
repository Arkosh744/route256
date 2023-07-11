package order_status

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type OrderStatusReceiver interface {
	Subscribe(topic string) error
}

type MessageSender interface {
	SendMessage(chatID int64, text string) error
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

func (r *receiver) Subscribe(topic string) error {
	partitionList, err := r.consumer.Partitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitionList {
		pc, err := r.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}

		go r.processMessages(pc, topic)
	}

	return nil
}

func (r *receiver) processMessages(pc sarama.PartitionConsumer, topic string) {
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

		if err := r.bot.SendMessage(r.chatID, messageText); err != nil {
			log.Printf("failed to send message: %v", err)
		}
	}
}
