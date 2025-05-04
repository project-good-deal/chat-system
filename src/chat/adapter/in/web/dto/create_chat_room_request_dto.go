package dto

import "github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in/command"

// ChatController handles chat-related requests
type CreateChatRoomRequestDto struct {
	UserId   string `json:"user_id" binding:"required"`
	DealId       string `json:"deal_id" binding:"required"`
	ChatRoomType string `json:"chat_room_type" binding:"required"`
}

func (c CreateChatRoomRequestDto) ToCommand() command.CreateChatRoomCommand {
	return command.CreateChatRoomCommand{
		UserId:   c.UserId,
		DealId:       c.DealId,
		ChatRoomType: c.ChatRoomType,
	}
}
