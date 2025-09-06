package consumer

import (
	"context"

	"github.com/en7ka/notifier/internal/client/kafka"
	"github.com/en7ka/notifier/internal/config"
)

type service struct {
	consumer   kafka.Consumer
	senderConf config.Sender
}

func NewService(
	consumer kafka.Consumer,
	senderConf config.Sender,
) *service {
	return &service{
		consumer:   consumer,
		senderConf: senderConf,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-s.run(ctx):
			if err != nil {
				return err
			}
		}
	}
}

func (s *service) run(ctx context.Context) <-chan error {
	errChan := make(chan error)

	go func() {
		defer close(errChan)

		errChan <- s.consumer.Consume(ctx, s.NoteSaveHandler)
	}()

	return errChan
}
