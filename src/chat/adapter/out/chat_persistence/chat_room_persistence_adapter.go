package chat_persistence

import (
	"fmt"
	"log"

	"github.com/project-good-deal/chat-system/src/chat/adapter/out/chat_persistence/dao"
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_out"
	"github.com/project-good-deal/chat-system/src/chat/domains"
	"github.com/project-good-deal/chat-system/src/util"
	"go.mongodb.org/mongo-driver/bson"
)

type ChatRoomPersistenceAdapter struct {
	// Add any dependencies you need here
}

func NewChatRoomPersistenceAdapter() chat_port_out.ChatRoomRepository {
	return &ChatRoomPersistenceAdapter{}
}

// Create creates a new chat message in the database
func (ChatRoomPersistenceAdapter) Create(chatRoom domains.ChatRoom) (domains.ChatRoom, error) {
	log.Printf("ChatPersistenceAdapter: Create chat with ID: %s", chatRoom.ChatRoomID)
	dbConnection, ctx, cancel := util.GetConnection()
  collection := dbConnection.Collection("chat_rooms")

  needToSave := dao.FromChatRoomEntity(chatRoom)
	res, err := collection.InsertOne(ctx, needToSave)
	if err != nil {
		fmt.Printf("Error inserting chat: %v\n", err)
		panic(err)
	}
	fmt.Printf("Inserted ID: %v\n", res.InsertedID)

	cancel()
	return chatRoom, nil
}

func (ChatRoomPersistenceAdapter) FindById(id string) domains.ChatRoom {
	log.Printf("ChatPersistenceAdapter: Find chat with chatRoom ID: %s", id)
	dbConnection, ctx, cancel := util.GetConnection()
	collection := dbConnection.Collection("chat_rooms")
	var chatRoomDao dao.ChatRoomDao
	err := collection.FindOne(ctx, bson.M{"chat_room_id": id}).Decode(&chatRoomDao)
	if err != nil {
		fmt.Printf("Error finding chat: %v\n", err)
		panic(err)
	}
	log.Printf("ChatPersistenceAdapter: Found chat room with ID: %s", chatRoomDao.ChatRoomID)
	defer cancel()
	return chatRoomDao.ToChatRoomEntity()
}

func (ChatRoomPersistenceAdapter) FindByUserId(userId string) []domains.ChatRoom {
	log.Printf("ChatPersistenceAdapter: Find chat with user ID: %s", userId)
	dbConnection, ctx, cancel := util.GetConnection()
	collection := dbConnection.Collection("chat_rooms")

	cursor, err := collection.Find(ctx, bson.M{"participant_ids": userId})
	if err != nil {
		fmt.Printf("Error finding chat: %v\n", err)
		panic(err)
	}

	var chatRooms []domains.ChatRoom
	for cursor.Next(ctx) {
		var chatRoomDao dao.ChatRoomDao
		if err := cursor.Decode(&chatRoomDao); err != nil {
			fmt.Printf("Error decoding chat room: %v\n", err)
			panic(err)
		}
		log.Printf("ChatPersistenceAdapter: Found chat room with ID: %s", chatRoomDao.ChatRoomID)
		chatRooms = append(chatRooms, chatRoomDao.ToChatRoomEntity())
	}

	defer cursor.Close(ctx)
	cancel()
	return chatRooms
}

func (ChatRoomPersistenceAdapter) UpdateChat(chatRoom domains.ChatRoom, chat domains.Chat) error {
	log.Printf("ChatPersistenceAdapter: Update chat with ID: %s", chatRoom.ChatRoomID)
	chatRoomDao := dao.FromChatRoomEntity(chatRoom)
	chatDao := dao.FromChatEntity(chat)

	dbConnection, ctx, cancel := util.GetConnection()
	collection := dbConnection.Collection("chat_rooms")

	filter := bson.M{"chat_room_id": chatRoomDao.ChatRoomID}
	update := bson.M{
		"$addToSet": bson.M{"chats": chatDao}, // 중복 없이 추가
		// "$push": bson.M{"user_ids": newUserId},  // 중복 허용
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Error updating chat room: %v", err)
		return err
	}

	defer cancel()
	return nil
}