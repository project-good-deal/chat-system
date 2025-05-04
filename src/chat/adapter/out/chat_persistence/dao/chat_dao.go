package dao

import (
	"github.com/project-good-deal/chat-system/src/chat/domains"
)

type ChatDao struct {
	ChatID  string `bson:"chat_id"`
	UserID  string `bson:"user_id"`
	Message string `bson:"message"`
}

func FromEntity(chat domains.Chat) *ChatDao {
	return &ChatDao{
		ChatID:  chat.ChatRoomID,
		UserID:  chat.UserID,
		Message: chat.Message,
	}
}