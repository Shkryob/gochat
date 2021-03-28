package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/shkryob/gochat/router/middleware"
	"github.com/shkryob/gochat/utils"
)

func (handler *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/users")
	guestUsers.POST("/signup", handler.SignUp)
	guestUsers.POST("/login", handler.Login)

	user := v1.Group("/users", jwtMiddleware)
	user.GET("", handler.GetUsers)
	user.GET("/me", handler.CurrentUser)
	user.GET("/:user_id", handler.GetUser)
	user.GET("/:user_id/avatar", handler.GetAvatar)
	user.POST("/:user_id/blacklist", handler.AddToBlackList)
	user.DELETE("/:user_id/blacklist", handler.RemoveFromBlackList)
	user.POST("/:user_id/friend", handler.AddFriend)
	user.DELETE("/:user_id/friend", handler.RemoveFriend)
	user.POST("/avatar", handler.UploadAvatar)

	chats := v1.Group("/chats")

	chats.GET("", handler.GetChats, jwtMiddleware)
	chats.GET("/:chat_id", handler.GetChat, jwtMiddleware)
	chats.POST("", handler.CreateChat, jwtMiddleware)
	chats.PUT("/:chat_id", handler.UpdateChat, jwtMiddleware)
	chats.DELETE("/:chat_id", handler.DeleteChat, jwtMiddleware)

	messages := v1.Group("/chats/:chat_id/messages")

	messages.GET("", handler.GetMessages, jwtMiddleware)
	messages.POST("", handler.AddMessage, jwtMiddleware)
	messages.DELETE("/:message_id", handler.DeleteMessage, jwtMiddleware)
	messages.PUT("/:message_id", handler.UpdateMessage, jwtMiddleware)

	sockets := v1.Group("/sockets")
	sockets.GET("", handler.GetSocket)
}
