package command

// ChatController handles chat-related requests
type CreateChatRoomCommand struct {
	UserId string
	DealId string
	ChatRoomType string
}