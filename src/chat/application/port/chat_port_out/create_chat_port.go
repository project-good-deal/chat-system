package chat_port_out

import (
	"github.com/project-good-deal/chat-system/src/chat/domains"
)

// ChatRepository is the interface that wraps the basic Create method
type ChatRepository interface {
	Create(chat *domains.Chat) (*domains.Chat, error)
}