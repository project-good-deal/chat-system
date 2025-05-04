package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-good-deal/chat-system/src/chat/adapter/in/web/dto"
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in"
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in/command"
)

type ChatRoomController struct {
	ChatRoomUseCase chat_port_in.ChatRoomUseCase
}

func NewChatRoomController(chatRoomUseCase chat_port_in.ChatRoomUseCase) *ChatRoomController {
	return &ChatRoomController{ChatRoomUseCase: chatRoomUseCase}
}

func (c *ChatRoomController) CreateChatRoom(ctx *gin.Context) {
	var chatRoomRequestDto dto.CreateChatRoomRequestDto
	if err := ctx.ShouldBindJSON(&chatRoomRequestDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chatRoomId, err := c.ChatRoomUseCase.CreateChatRoom(chatRoomRequestDto.ToCommand())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Chat room created with ID: %s", chatRoomId)
	ctx.JSON(http.StatusOK, gin.H{"chat_room_id": chatRoomId})
}

// GetChatRoom retrieves a chat room by its ID
func (c *ChatRoomController) GetChatRoomByUserId(ctx *gin.Context) {
	userId := ctx.Query("user-id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user-id is required"})
		return
	}

	chatRooms := c.ChatRoomUseCase.GetMyChatRooms(command.ReadMyChatRoomsCommand{
		UserId: userId,
  })

	ctx.JSON(http.StatusOK, chatRooms)
}