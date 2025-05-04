package chat_service

import (
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in"
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in/command"
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_out"
	"github.com/project-good-deal/chat-system/src/chat/domains"
)

// ChatRoomService is a struct that implements the ChatUseCase interface
type ChatRoomService struct {
	chatRoomRepository chat_port_out.ChatRoomRepository
}

// NewChatRoomService creates a new instance of ChatRoomService
func NewChatRoomService(chatRoomRepository chat_port_out.ChatRoomRepository) chat_port_in.ChatRoomUseCase {
	return ChatRoomService{
		chatRoomRepository: chatRoomRepository,
	}
}

// CreateChatRoom creates a new chat room
func (service ChatRoomService) CreateChatRoom(command command.CreateChatRoomCommand) (string, error) {
	// Create a new chat room
	chatRoom := domains.CreateChatRoom(command.DealId, []string{command.UserId}, command.ChatRoomType)

	// Save the chat room to the repository
	_, err := service.chatRoomRepository.Create(chatRoom)
	if err != nil {
		return "", err
	}

	return chatRoom.ChatRoomID, nil
}

// CreateChatRoom creates a new chat room
func (service ChatRoomService) GetMyChatRooms(command command.ReadMyChatRoomsCommand) []domains.ChatRoom {
	// Save the chat room to the repository
	chatrooms := service.chatRoomRepository.FindByUserId(command.UserId)
	return chatrooms
}