package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	tokenEnvName = "BOT_TOKEN"
	idEnvName    = "ID"
)

type Sender interface {
	Token() string
	ID() int64
}

type SenderConfig struct {
	token  string
	chatID string
}

func NewSenderConfig() (Sender, error) {
	token := os.Getenv(tokenEnvName)
	if token == "" {
		return nil, fmt.Errorf("token not found in environment variable %s", tokenEnvName)
	}

	id := os.Getenv(idEnvName)
	if id == "" {
		return nil, fmt.Errorf("id not found in environment variable %s", idEnvName)
	}

	return &SenderConfig{
		token:  token,
		chatID: id,
	}, nil
}

func (s *SenderConfig) Token() string {
	return s.token
}

func (s *SenderConfig) ID() int64 {
	intID, err := strconv.Atoi(s.chatID)
	if err != nil {
		log.Fatalf("error converting id to int: %v", err)
	}
	return int64(intID)
}
