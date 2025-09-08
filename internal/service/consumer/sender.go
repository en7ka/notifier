package consumer

import (
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *service) SendToTelegram(text string) error {
	bot, err := tgbotapi.NewBotAPI(s.senderConf.Token())
	if err != nil {
		return fmt.Errorf("new telegram bot: %w", err)
	}

	msg := tgbotapi.NewMessage(s.senderConf.ID(), text)

	log.Printf("--- [DEBUG] Attempting to send message to Telegram chat ID: %d ---", msg)

	_, err = bot.Send(msg)

	if err != nil {
		log.Printf("Error sending message: %v", err)
		return fmt.Errorf("send message: %w", err)
	}

	log.Printf("--- [SUCCESS] Message sent to Telegram successfully! ---")

	return nil
}
