package domains

import (
	"time"

	"github.com/google/uuid"

	"github.com/project-good-deal/chat-system/src/chat/domains/types"
)

type ChatRoom struct {
	ChatRoomID     string             `json:"chat_room_id"`
	DealID         string             `json:"deal_id"`
	ParticipantIDs []string           `json:"participant_ids"`
	ChatRoomType   types.ChatRoomType `json:"type"`
	Chats          []Chat             `json:"chats"`
	CreatedAt  time.Time `json:"created_at"`
}

func (c ChatRoom) AddChat(chat Chat) {
	c.Chats = append(c.Chats, chat)
}

func CreateChatRoom(dealId string, userIds []string, chatRoomType string) ChatRoom {
	createdAt := time.Now()
	return ChatRoom{
		ChatRoomID:     uuid.NewString(),
		ParticipantIDs: userIds,
		Chats:          []Chat{},
		DealID:         dealId,
		ChatRoomType:   types.GetByString(chatRoomType),
		CreatedAt:      createdAt,
	}
}
