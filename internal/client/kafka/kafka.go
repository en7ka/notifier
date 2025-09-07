package kafka

import (
	"context"

	"github.com/en7ka/notifier/internal/client/kafka/consumer"
	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	Consume(ctx context.Context, handler consumer.Handler) error
	Close() error
}

type Message = kafka.Message
