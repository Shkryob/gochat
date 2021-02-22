package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shkryob/gochat/model"
	"github.com/shkryob/gochat/utils"
)

type chatResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Admin  struct {
		Username string `json:"username"`
	} `json:"admin"`
}

type singleChatResponse struct {
	Chat *chatResponse `json:"chat"`
}

type chatListResponse struct {
	Chats      []*chatResponse `json:"chats"`
	ChatsCount int             `json:"chatsCount"`
}

func newChatResponse(c echo.Context, ch *model.Chat) *singleChatResponse {
	chat := new(chatResponse)
	chat.ID = ch.ID
	chat.Title = ch.Title
	chat.Admin.Username = ch.Admin.Username
	return &singleChatResponse{chat}
}

func newChatListResponse(chats []model.Chat, count int) *chatListResponse {
	r := new(chatListResponse)
	r.Chats = make([]*chatResponse, 0)
	for _, c := range chats {
		chat := new(chatResponse)
		chat.ID = c.ID
		chat.Title = c.Title
		chat.Admin.Username = c.Admin.Username
		r.Chats = append(r.Chats, chat)
	}
	r.ChatsCount = count
	return r
}

type messageResponse struct {
	ID        uint      `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      struct {
		Username string `json:"username"`
	} `json:"user"`
}

type singleMessageResponse struct {
	Message *messageResponse `json:"message"`
}

type messageListResponse struct {
	Messages []messageResponse `json:"messages"`
}

func newMessageResponse(c echo.Context, m *model.Message) *singleMessageResponse {
	message := new(messageResponse)
	message.ID = m.ID
	message.Body = m.Body
	message.CreatedAt = m.CreatedAt
	message.UpdatedAt = m.UpdatedAt
	message.User.Username = m.User.Username
	return &singleMessageResponse{message}
}

func newMessageListResponse(c echo.Context, messages []model.Message) *messageListResponse {
	r := new(messageListResponse)
	mr := messageResponse{}
	r.Messages = make([]messageResponse, 0)
	for _, i := range messages {
		mr.ID = i.ID
		mr.Body = i.Body
		mr.CreatedAt = i.CreatedAt
		mr.UpdatedAt = i.UpdatedAt
		mr.User.Username = i.User.Username

		r.Messages = append(r.Messages, mr)
	}
	return r
}

type userResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}
