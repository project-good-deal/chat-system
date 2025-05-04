package chat_port_in

import (
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in/command"
	"github.com/project-good-deal/chat-system/src/chat/domains"
)

type ChatRoomUseCase interface {
	CreateChatRoom(command.CreateChatRoomCommand) (string, error)
	GetMyChatRooms(command.ReadMyChatRoomsCommand) []domains.ChatRoom
}