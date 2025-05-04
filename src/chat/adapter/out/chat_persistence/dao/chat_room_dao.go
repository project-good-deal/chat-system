package dao

import (
	"time"

	"github.com/project-good-deal/chat-system/src/chat/domains"
	"github.com/project-good-deal/chat-system/src/chat/domains/types"
)

type ChatRoomDao struct {
	ChatRoomID     string             `bson:"chat_room_id"`
	DealID         string             `bson:"deal_id"`
	ParticipantIDs []string           `bson:"participant_ids"`
	Chats          []ChatDao          `bson:"chats"`
	ChatRoomType   types.ChatRoomType `bson:"chat_room_type"`
	CreatedAt      time.Time             `bson:"created_at"`
}

func (c ChatRoomDao) ToChatRoomEntity() domains.ChatRoom {
	chats := mapChatDaosToChat(c.Chats)

	return domains.ChatRoom{
		ChatRoomID:     c.ChatRoomID,
		ParticipantIDs: c.ParticipantIDs,
		DealID:         c.DealID,
		Chats:          chats,
		ChatRoomType:   c.ChatRoomType,
		CreatedAt:      c.CreatedAt,
	}
}

func FromChatRoomEntity(chatRoom domains.ChatRoom) *ChatRoomDao {
	chatDaos := mapChatsToChatDaos(chatRoom.Chats)

	return &ChatRoomDao{
		ChatRoomID:     chatRoom.ChatRoomID,
		ParticipantIDs: chatRoom.ParticipantIDs,
		DealID:         chatRoom.DealID,
		Chats:           chatDaos,
		ChatRoomType:   chatRoom.ChatRoomType,
		CreatedAt:      chatRoom.CreatedAt,
	}
}

func mapChatsToChatDaos(chats []domains.Chat) []ChatDao {
	var result []ChatDao = make([]ChatDao, 0, len(chats))
	for _, chat := range chats {
		result = append(result, *FromChatEntity(chat))
	}
	return result
}

func mapChatDaosToChat(daos []ChatDao) []domains.Chat {
	var result []domains.Chat = make([]domains.Chat, 0, len(daos))
	for _, dao := range daos {
		result = append(result, dao.ToChatEntity())
	}
	return result
}