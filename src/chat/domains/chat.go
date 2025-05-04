package domains

import "time"

type Chat struct {
	// ID         int `json:"id"`
	ChatRoomID     string `json:"chat_room_id"`
	UserID     string `json:"user_id"`
	Message    string `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}

func CreateChat(chatRoomId string, userId string, message string) Chat {
	createdAt := time.Now()
	return Chat{
		ChatRoomID: chatRoomId,
		UserID: userId,
		Message: message,
		CreatedAt: createdAt,
	}
}
