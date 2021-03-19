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
		ID uint `json:"id"`
		Username string `json:"username"`
	} `json:"admin"`
	Participants      []*simplifiedUserResponse `json:"participants"`
}

type singleChatResponse struct {
	Chat *chatResponse `json:"chat"`
}

type chatListResponse struct {
	Chats      []*chatResponse `json:"chats"`
	ChatsCount int             `json:"chatsCount"`
}

type userListResponse struct {
	Users      []*simplifiedUserResponse `json:"users"`
	UsersCount int             `json:"usersCount"`
}

func newChatResponse(c echo.Context, ch *model.Chat) *singleChatResponse {
	chat := new(chatResponse)
	chat.ID = ch.ID
	chat.Title = ch.Title
	chat.Admin.Username = ch.Admin.Username
	chat.Admin.ID = ch.Admin.ID
	chat.Participants = make([]*simplifiedUserResponse, 0)
	for _, c := range ch.Users {
		participant := new(simplifiedUserResponse)
		participant.ID = c.ID
		participant.Username = c.Username
		chat.Participants = append(chat.Participants, participant)
	}
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
		chat.Admin.ID = c.Admin.ID
		chat.Participants = make([]*simplifiedUserResponse, 0)
		for _, ch := range c.Users {
			participant := new(simplifiedUserResponse)
			participant.ID = ch.ID
			participant.Username = ch.Username
			chat.Participants = append(chat.Participants, participant)
		}
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
		ID uint `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
	Chat      struct {
		ID uint `json:"id"`
	} `json:"chat"`
}

type SingleMessageResponse struct {
	Message *messageResponse `json:"message"`
}

type messageListResponse struct {
	Messages []messageResponse `json:"messages"`
}

func newMessageResponse(c echo.Context, m *model.Message) *SingleMessageResponse {
	message := new(messageResponse)
	message.ID = m.ID
	message.Body = m.Body
	message.CreatedAt = m.CreatedAt
	message.UpdatedAt = m.UpdatedAt
	message.User.ID = m.User.ID
	message.User.Username = m.User.Username
	message.Chat.ID = m.ChatID
	return &SingleMessageResponse{message}
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
		mr.User.ID = i.User.ID
		mr.User.Username = i.User.Username

		r.Messages = append(r.Messages, mr)
	}
	return r
}

type userResponse struct {
	User struct {
		ID        uint      `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	} `json:"user"`
}

type simplifiedUserResponse struct {
	ID   	 uint   `json:"id"`
	Username string `json:"username"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.ID = u.ID
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}

func newSimplifiedUserResponse(u *model.User) *simplifiedUserResponse {
	r := new(simplifiedUserResponse)
	r.ID = u.ID
	r.Username = u.Username
	return r
}

func newUserListResponse(users []model.User, count int) *userListResponse {
	r := new(userListResponse)
	r.Users = make([]*simplifiedUserResponse, 0)
	for _, u := range users {
		user := new(simplifiedUserResponse)

		user.Username = u.Username
		user.ID = u.ID
		r.Users = append(r.Users, user)
	}
	r.UsersCount = count
	return r
}