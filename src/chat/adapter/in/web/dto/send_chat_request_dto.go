package dto

// ChatController handles chat-related requests
type SendChatRequestDto struct {
	ChatRoomId string `json:"chat_room_id" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}