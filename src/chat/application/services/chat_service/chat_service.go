package chat_service

import (
	// import my domain chat
	"log"

	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in"
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_out"
	"github.com/project-good-deal/chat-system/src/chat/domains"
)

// chatService implements the ChatUseCase interface
type ChatService struct {
	chatRepository chat_port_out.ChatRepository
	// Add any dependencies you need here
}

func NewChatService(chatRepository chat_port_out.ChatRepository) chat_port_in.ChatUseCase {
	return &ChatService{chatRepository: chatRepository}
}

func (chartService ChatService) Create(chatRoomId string, userId string, message string) (domains.Chat) {
	chat := domains.CreateChat(chatRoomId, userId, message)
	log.Printf("Chat created: %v", chat.ChatRoomID)
	log.Printf("Chat created: %v", chat.UserID)
	log.Printf("Chat created: %v", chat.Message)
	
  chartService.chatRepository.Create(chat)
	
	return chat
}