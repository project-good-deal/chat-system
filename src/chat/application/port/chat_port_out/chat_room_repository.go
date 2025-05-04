package chat_port_out

import (
	"github.com/project-good-deal/chat-system/src/chat/domains"
)

// ChatRepository is the interface that wraps the basic Create method
type ChatRoomRepository interface {
	Create(chatRoom domains.ChatRoom) (domains.ChatRoom, error)
	FindByUserId(userId string) []domains.ChatRoom
	FindById(id string) domains.ChatRoom
	UpdateChat(chatRoom domains.ChatRoom, chat domains.Chat) error
}
