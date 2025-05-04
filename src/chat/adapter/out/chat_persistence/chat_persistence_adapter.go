package chat_persistence

import (
	"fmt"
	"log"

	"github.com/project-good-deal/chat-system/src/chat/adapter/out/chat_persistence/dao"
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_out"
	"github.com/project-good-deal/chat-system/src/chat/domains"
	"github.com/project-good-deal/chat-system/src/util"
)

type ChatPersistenceAdapter struct {
	// Add any dependencies you need here
}

func NewChatPersistenceAdapter() chat_port_out.ChatRepository {
	return &ChatPersistenceAdapter{}
}

// Create a new chat message in the database
func (ChatPersistenceAdapter) Create(chat domains.Chat) (domains.Chat, error) {
	log.Printf("ChatPersistenceAdapter: Create chat with ID: %s", chat.ChatRoomID)
	dbConnection, ctx, cancel := util.GetConnection()
  collection := dbConnection.Collection("chats")

  needToSave := dao.FromChatEntity(chat)
	res, err := collection.InsertOne(ctx, needToSave)
	if err != nil {
		fmt.Printf("Error inserting chat: %v\n", err)
		panic(err)
	}
	fmt.Printf("Inserted ID: %v\n", res.InsertedID)

	cancel()
	return chat, nil
}