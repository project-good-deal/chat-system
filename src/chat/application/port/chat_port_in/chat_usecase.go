package chat_port_in

import (
	"github.com/project-good-deal/chat-system/src/chat/domains"
)

type ChatUseCase interface {
	Create(chatRoomId string, userId string, message string) (domains.Chat)
}
