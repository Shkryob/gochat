package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/shkryob/gochat/router/middleware"
	"github.com/shkryob/gochat/utils"
)

func (handler *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/users")
	guestUsers.POST("", handler.SignUp)
	guestUsers.POST("/login", handler.Login)

	user := v1.Group("/users", jwtMiddleware)
	user.GET("", handler.GetUsers)
	user.GET("/me", handler.CurrentUser)

	chats := v1.Group("/chats")

	chats.GET("", handler.GetChats, jwtMiddleware)
	chats.GET("/:chat_id", handler.GetChat, jwtMiddleware)
	chats.POST("", handler.CreateChat, jwtMiddleware)
	chats.PUT("/:chat_id", handler.UpdateChat, jwtMiddleware)
	chats.DELETE("/:chat_id", handler.DeleteChat, jwtMiddleware)

	messages := v1.Group("/chats/:chat_id/messages")

	messages.POST("", handler.AddMessage, jwtMiddleware)
	messages.DELETE("/:message_id", handler.DeleteMessage, jwtMiddleware)
	messages.PUT("/:message_id", handler.UpdateMessage, jwtMiddleware)
}
