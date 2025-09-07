package config

import (
	"fmt"
	"os"
	"strings"
)

const (
	brokersEnvName = "KAFKA_BROKERS"
	groupEnvName   = "KAFKA_GROUP"
	topicEnvName   = "THEME"
)

type KafkaConsumerConfig interface {
	Brokers() []string
	GroupID() string
	Topic() string
}

type kafkaConfig struct {
	brokers []string
	groupID string
	topic   string
}

func NewKafkaConfig() (KafkaConsumerConfig, error) {
	brokersRaw := os.Getenv(brokersEnvName)
	groupID := os.Getenv(groupEnvName)
	topic := os.Getenv(topicEnvName)

	if brokersRaw == "" {
		return nil, fmt.Errorf("%s is empty", brokersEnvName)
	}
	if groupID == "" {
		return nil, fmt.Errorf("%s is empty", groupEnvName)
	}
	if topic == "" {
		return nil, fmt.Errorf("%s is empty", topicEnvName)
	}

	return &kafkaConfig{
		brokers: splitCSV(brokersRaw),
		groupID: groupID,
		topic:   topic,
	}, nil
}

func (k *kafkaConfig) Brokers() []string { return k.brokers }
func (k *kafkaConfig) GroupID() string   { return k.groupID }
func (k *kafkaConfig) Topic() string     { return k.topic }

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
