package app

import (
	"log"

	kafkaiface "github.com/en7ka/notifier/internal/client/kafka"
	consumerc "github.com/en7ka/notifier/internal/client/kafka/consumer"

	"github.com/en7ka/notifier/internal/closer"
	"github.com/en7ka/notifier/internal/config"

	svciface "github.com/en7ka/notifier/internal/service"
	svcsender "github.com/en7ka/notifier/internal/service/consumer"
)

type serviceProvider struct {
	kafkaConsumerConfig config.KafkaConsumerConfig
	senderConfig        config.Sender

	senderConsumer svciface.ConsumerService

	consumer kafkaiface.Consumer
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) KafkaConsumerConfig() config.KafkaConsumerConfig {
	if s.kafkaConsumerConfig == nil {
		cfg, err := config.NewKafkaConfig()
		if err != nil {
			log.Fatalf("Error creating kafka consumer config: %v", err)
		}
		s.kafkaConsumerConfig = cfg
	}
	return s.kafkaConsumerConfig
}

func (s *serviceProvider) SenderConfig() config.Sender {
	if s.senderConfig == nil {
		cfg, err := config.NewSenderConfig()
		if err != nil {
			log.Fatalf("Error creating sender config: %v", err)
		}
		s.senderConfig = cfg
	}
	return s.senderConfig
}

func (s *serviceProvider) Consumer() kafkaiface.Consumer {
	if s.consumer == nil {
		cfg := s.KafkaConsumerConfig()

		s.consumer = consumerc.New(
			cfg.Brokers(),
			cfg.GroupID(),
			cfg.Topic(),
		)
		closer.Add(s.consumer.Close)
	}
	return s.consumer
}

func (s *serviceProvider) SenderConsumer() svciface.ConsumerService {
	if s.senderConsumer == nil {
		s.senderConsumer = svcsender.NewService(
			s.Consumer(),
			s.SenderConfig(),
		)
	}
	return s.senderConsumer
}
