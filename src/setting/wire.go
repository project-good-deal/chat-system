//go:build wireinject
// +build wireinject

package setting

import (
	"github.com/google/wire"
	"github.com/project-good-deal/chat-system/src/chat/adapter/in/web"
	"github.com/project-good-deal/chat-system/src/chat/adapter/out/chat_persistence"
	"github.com/project-good-deal/chat-system/src/chat/application/services/chat_service"
)

func InitializeChatController() *web.ChatController {
	// Initialize the chat service and controller
	// using wire to inject dependencies
	wire.Build(web.NewChatController, chat_service.NewChatService, chat_persistence.NewChatRoomPersistenceAdapter)
	return &web.ChatController{}
}

func InitializeChatRoomController() *web.ChatRoomController {
	// Initialize the chat service and controller
	// using wire to inject dependencies
	wire.Build(web.NewChatRoomController, chat_service.NewChatRoomService, chat_persistence.NewChatRoomPersistenceAdapter)
	return &web.ChatRoomController{}
}