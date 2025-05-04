package dao

import (
	"time"

	"github.com/project-good-deal/chat-system/src/chat/domains"
)

type ChatDao struct {
	ChatID  string `bson:"chat_id"`
	UserID  string `bson:"user_id"`
	Message string `bson:"message"`
	CreatedAt time.Time `bson:"created_at"`
}

func FromChatEntity(chat domains.Chat) *ChatDao {
	return &ChatDao{
		ChatID:  chat.ChatRoomID,
		UserID:  chat.UserID,
		Message: chat.Message,
		CreatedAt: chat.CreatedAt,
	}
}

func (dao ChatDao) ToChatEntity() domains.Chat {
	return domains.Chat{
		UserID:         dao.UserID,
		Message:        dao.Message,
		CreatedAt:     dao.CreatedAt,
	}
}