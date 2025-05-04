package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-good-deal/chat-system/src/chat/adapter/in/web/dto"
	"github.com/project-good-deal/chat-system/src/chat/application/port/chat_port_in"
)

type ChatController struct {
	usecase chat_port_in.ChatUseCase
}

func NewChatController(usecase chat_port_in.ChatUseCase) *ChatController {
	return &ChatController{usecase: usecase}
}

// CreateChat handles the creation of a new chat message
// @Summary Create a new chat message
// @Description Create a new chat message
// @Tags chat
// @Accept json
// @Produce json
// @Param chat body ChatRequestDto true "Chat message"
// @Success 200 {object} domains.Chat
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chat [post]
func (c ChatController) CreateChat(ctx *gin.Context) {
	var chatRequestDto dto.SendChatRequestDto
	log.Printf("hihi")
	if err := ctx.ShouldBindJSON(&chatRequestDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat := c.usecase.Create(chatRequestDto.ChatRoomId, chatRequestDto.UserId, chatRequestDto.Message)
	ctx.JSON(http.StatusOK, chat)
}