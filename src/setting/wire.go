//go:build wireinject
// +build wireinject

package setting

import (
	"github.com/google/wire"
	"github.com/project-good-deal/chat-system/src/chat/adapter/in/web"
	"github.com/project-good-deal/chat-system/src/chat/application/services/chat_service"
)

func InitializeChatService() *web.ChatController {
	// Initialize the chat service and controller
	// using wire to inject dependencies
	wire.Build(web.NewChatController, chat_service.NewChatService)
	return &web.ChatController{}
}