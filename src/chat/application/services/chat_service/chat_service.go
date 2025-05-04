package chat_service

import (
	// import my domain chat

	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in"
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_out"
	"github.com/project-good-deal/chat-system/src/chat/domains"
)

// chatService implements the ChatUseCase interface
type ChatService struct {
	chatRoomRepository chat_port_out.ChatRoomRepository
	// chatRepository chat_port_out.ChatRepository
	// Add any dependencies you need here
}

func NewChatService(chatRoomRepository chat_port_out.ChatRoomRepository) chat_port_in.ChatUseCase {
	return &ChatService{chatRoomRepository: chatRoomRepository}
}

func (chartService ChatService) Create(chatRoomId string, userId string, message string) (domains.Chat) {
	chat := domains.CreateChat(chatRoomId, userId, message)
	
  chatRoom := chartService.chatRoomRepository.FindById(chatRoomId)
	chatRoom.AddChat(chat)
	
	// Save the chat to the repository
	chartService.chatRoomRepository.UpdateChat(chatRoom, chat)

	return chat
}