package chat_persistence

import (
	"log"

	"github.com/project-good-deal/chat-system/src/chat/domains"
)

// Create creates a new chat message in the database
func Create(chat *domains.Chat) (*domains.Chat, error) {
	log.Printf("Chat created: %v", chat.ID)
	return chat, nil
}