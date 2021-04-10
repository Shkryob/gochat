package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/shkryob/gochat/model"
	"github.com/shkryob/gochat/utils"
)

type userRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (request *userRegisterRequest) bind(context echo.Context, user *model.User) error {
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	user.Username = request.User.Username
	user.Email = request.User.Email
	h, err := utils.HashPassword(request.User.Password)
	if err != nil {
		return err
	}
	user.Password = h
	return nil
}

type userLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type chatCreateRequest struct {
	Chat struct {
		Title string `json:"title"`
		Participants []uint
	} `json:"chat"`
}

func (request *chatCreateRequest) bind(context echo.Context, a *model.Chat) error {
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	a.Title = request.Chat.Title
	return nil
}

type chatUpdateRequest struct {
	Chat struct {
		Title string `json:"title"`
	} `json:"chat"`
}

func (request *chatUpdateRequest) populate(chat *model.Chat) {
	request.Chat.Title = chat.Title
}

func (request *chatUpdateRequest) bind(context echo.Context, a *model.Chat) error {
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	a.Title = request.Chat.Title
	return nil
}

type createMessageRequest struct {
	Message struct {
		Body string `json:"body" validate:"required"`
	} `json:"message"`
}


type updateMessageRequest struct {
	Message struct {
		Body string `json:"body" validate:"required"`
	} `json:"message"`
}

func (request *createMessageRequest) bind(context echo.Context, cm *model.Message) error {
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	cm.Body = request.Message.Body
	return nil
}

func (request *updateMessageRequest) bind(context echo.Context, cm *model.Message) error {
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	cm.Body = request.Message.Body
	return nil
}