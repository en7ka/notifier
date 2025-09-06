package consumer

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Handler func(ctx context.Context, msg kafka.Message) error

type Consumer struct {
	r *kafka.Reader
}

func New(brokers []string, groupID, topic string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         groupID,
		Topic:           topic,
		StartOffset:     kafka.LastOffset,
		MinBytes:        1,
		MaxBytes:        10e6,
		MaxWait:         500 * time.Millisecond,
		ReadLagInterval: 5 * time.Second,
	})
	return &Consumer{r: r}
}

func (c *Consumer) Consume(ctx context.Context, handler Handler) error {
	for {
		m, err := c.r.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			log.Println(err)
			continue
		}
		if err = handler(ctx, m); err != nil {
			log.Printf("handler error: %v", err)
		}
	}
}

func (c *Consumer) Close() error {
	log.Println("Closing kafka consumer reader")
	return c.r.Close()
}
