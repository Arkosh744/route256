package status

import (
	"fmt"
	"time"

	"route256/libs/log"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type OrderStatusSender struct {
	producer sarama.SyncProducer
	topic    string
}

type Handler func(id string)

func NewOrderStatusSender(producer sarama.SyncProducer, topic string) OrderStatusSender {
	return OrderStatusSender{
		producer: producer,
		topic:    topic,
	}
}

func (o OrderStatusSender) SendOrderStatus(orderID int64, status string) error {
	msg := &sarama.ProducerMessage{
		Topic:     o.topic,
		Key:       sarama.StringEncoder(fmt.Sprint(orderID)),
		Value:     sarama.StringEncoder(status),
		Partition: -1,
		Timestamp: time.Now(),
	}

	partition, offset, err := o.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Info("notification sent",
		zap.Int64("orderID", orderID),
		zap.String("status", status),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
	)

	return nil
}
