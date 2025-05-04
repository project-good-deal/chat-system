package chat_service

import (
	// import my domain chat
	"log"

	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in"
	"github.com/project-good-deal/chat-system/src/chat/domains"
)

// chatService implements the ChatUseCase interface
type ChatService struct {
	// Add any dependencies you need here
}

func NewChatService() chat_port_in.ChatUseCase {
	return &ChatService{}
}

func (s *ChatService) Create(userId string, message string) (*domains.Chat) {
	chat := domains.CreateChat(userId, message)
	log.Printf("Chat created: %v", chat.ID)
	log.Printf("Chat created: %v", chat.UserID)
	log.Printf("Chat created: %v", chat.Message)
	return chat
}