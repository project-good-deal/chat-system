package domains

import (
	"github.com/jinzhu/gorm"
)

type Chat struct {
	Model			gorm.Model

	ID         int `json:"id"`
	UserID     string `json:"user_id"`
	Message    string `json:"message"`
	CreatedAt  string `json:"created_at"`
}

func CreateChat(userID string, message string) *Chat {
	return &Chat{
		ID: 0,
		UserID: "hello",
		Message: "hi",
	}
}
