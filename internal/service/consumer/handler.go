package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/en7ka/notifier/internal/model"
	"github.com/segmentio/kafka-go"
)

func (s *service) NoteSaveHandler(_ context.Context, msg kafka.Message) error {
	var user model.User

	log.Printf("[msg] topic=%s part=%d off=%d key=%s val=%s",
		msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

	if err := json.Unmarshal(msg.Value, &user); err != nil {
		return fmt.Errorf("ошибка декодирования: %w", err)
	}

	messageSend := fmt.Sprintf("Зарегался пользователь %s c ролью: %s", user.Name, user.Role)

	if err := s.SendToTelegram(messageSend); err != nil {
		log.Printf("Ошибка отправки в Telegram: %v", err)
		return err
	}

	log.Println("Сообщение успешно отправлено в Telegram.")
	return nil
}
