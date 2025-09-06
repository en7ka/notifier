package consumer

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *service) SendToTelegram(text string) error {
	bot, err := tgbotapi.NewBotAPI(s.senderConf.Token())
	if err != nil {
		return fmt.Errorf("new telegram bot: %w", err)
	}

	msg := tgbotapi.NewMessage(s.senderConf.ID(), text)
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
