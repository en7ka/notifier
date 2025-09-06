package consumer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type MsgHandler func(ctx context.Context, msg kafka.Message) error

type GroupMsgHandler struct {
	handler MsgHandler
}

func NewGroupMsgHandler(handler MsgHandler) *GroupMsgHandler {
	return &GroupMsgHandler{handler: handler}
}

// Run — запускает чтение сообщений с Reader и прокидывает их в msgHandler
func (h *GroupMsgHandler) Run(ctx context.Context, r *kafka.Reader) error {
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				log.Printf("Error closing reader: %v", err)
				return nil
			}
			log.Printf("Error closing reader: %v", err)
			continue
		}
		//просто лог для понимания, что за сообщение пришло
		log.Printf("got msg topic=%s part=%d off=%d val=%s\n",
			m.Topic, m.Partition, m.Offset, string(m.Value))

		if h.handler != nil {
			if err := h.handler(ctx, m); err != nil {
				log.Printf("Error closing reader: %v", err)
			}
		}
	}
}
