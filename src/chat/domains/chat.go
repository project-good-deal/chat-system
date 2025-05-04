package domains

import (
	"github.com/jinzhu/gorm"
)

type Chat struct {
	Model			gorm.Model

	// ID         int `json:"id"`
	ChatRoomID     string `json:"chat_room_id"`
	UserID     string `json:"user_id"`
	Message    string `json:"message"`
	// CreatedAt  string `json:"created_at"`
}

func CreateChat(chatRoomId string, userId string, message string) Chat {
	return Chat{
		ChatRoomID: chatRoomId,
		UserID: userId,
		Message: message,
	}
}
