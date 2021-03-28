package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/shkryob/gochat/model"
	"github.com/shkryob/gochat/utils"
)

func (handler *Handler) GetChat(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("chat_id"), 10, 32)
	id := uint(id64)
	chat, err := handler.chatStore.GetById(id)

	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	if chat == nil {
		return utils.ResponseByContentType(context, http.StatusNotFound, utils.NotFound())
	}

	return utils.ResponseByContentType(context, http.StatusOK, newChatResponse(context, chat, ""))
}

func (handler *Handler) GetChats(context echo.Context) error {
	var (
		chats []model.Chat
		count int
	)

	offset, err := strconv.Atoi(context.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(context.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	chats, count, err = handler.chatStore.List(offset, limit)

	return utils.ResponseByContentType(context, http.StatusOK, newChatListResponse(chats, count))
}

func (handler *Handler) CreateChat(context echo.Context) error {
	var chat model.Chat

	req := &chatCreateRequest{}
	if err := req.bind(context, &chat); err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}

	chat.AdminID = userIDFromToken(context)

	if chat.AdminID == 0 {
		return utils.ResponseByContentType(context, http.StatusUnauthorized, utils.NewError(errors.New("unauthorized action")))
	}

	if err := handler.ValidateParticipants(req.Chat.Participants); err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}

	err := handler.chatStore.CreateChat(&chat)

	req.Chat.Participants = append(req.Chat.Participants, chat.AdminID)
	participants := handler.userStore.GetByIDs(req.Chat.Participants)
	handler.chatStore.ReplaceParticipants(&chat, &participants)

	if err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}

	chatResponse := newChatResponse(context, &chat, "chat_created")
	BroadcastChatCreated(&chat, chatResponse)
	return utils.ResponseByContentType(context, http.StatusCreated, chatResponse)
}

func (handler *Handler) UpdateChat(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("chat_id"), 10, 32)
	id := uint(id64)

	chat, err := handler.chatStore.GetById(id)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	if chat == nil {
		return utils.ResponseByContentType(context, http.StatusNotFound, utils.NotFound())
	}

	req := &chatUpdateRequest{}
	req.populate(chat)

	if err := req.bind(context, chat); err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if chat.AdminID != userIDFromToken(context) {
		return utils.ResponseByContentType(context, http.StatusUnauthorized, utils.NewError(errors.New("unauthorized action")))
	}

	if err = handler.chatStore.UpdateChat(chat); err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}
	chatResponse := newChatResponse(context, chat, "chat_updated")

	return utils.ResponseByContentType(context, http.StatusOK, chatResponse)
}

func (handler *Handler) DeleteChat(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("chat_id"), 10, 32)
	id := uint(id64)

	chat, err := handler.chatStore.GetById(id)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	if chat == nil {
		return utils.ResponseByContentType(context, http.StatusNotFound, utils.NotFound())
	}

	if chat.AdminID != userIDFromToken(context) {
		return utils.ResponseByContentType(context, http.StatusUnauthorized, utils.NewError(errors.New("unauthorized action")))
	}

	err = handler.chatStore.DeleteChat(chat)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	return utils.ResponseByContentType(context, http.StatusOK, map[string]interface{}{"result": "ok"})
}

func (handler *Handler) AddMessage(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("chat_id"), 10, 32)
	id := uint(id64)

	chat, err := handler.chatStore.GetById(id)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	if chat == nil {
		return utils.ResponseByContentType(context, http.StatusNotFound, utils.NotFound())
	}

	var cm model.Message

	req := &createMessageRequest{}
	if err := req.bind(context, &cm); err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}

	cm.UserID = userIDFromToken(context)

	if cm.UserID == 0 {
		return utils.ResponseByContentType(context, http.StatusUnauthorized, utils.NewError(errors.New("unauthorized action")))
	}

	if err = handler.chatStore.AddMessage(chat, &cm); err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	response := newMessageResponse(context, &cm, "message_created")
	curUser, _ := handler.GetCurrentUser(context)
	BroadcastMessage(chat, response, curUser.BlacklistedBy)

	return utils.ResponseByContentType(context, http.StatusCreated, response)
}

func (handler *Handler) GetMessages(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("chat_id"), 10, 32)
	id := uint(id64)

	curUser, _ := handler.GetCurrentUser(context)
	cm, err := handler.chatStore.GetMessagesByChatId(id, curUser)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	return utils.ResponseByContentType(context, http.StatusOK, newMessageListResponse(context, cm))
}

func (handler *Handler) DeleteMessage(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("message_id"), 10, 32)
	id := uint(id64)

	if err != nil {
		return utils.ResponseByContentType(context, http.StatusBadRequest, utils.NewError(err))
	}

	cm, err := handler.chatStore.GetMessageByID(id)
	chat, err := handler.chatStore.GetById(cm.ChatID)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	if cm == nil {
		return utils.ResponseByContentType(context, http.StatusNotFound, utils.NotFound())
	}

	if cm.UserID != userIDFromToken(context) {
		return utils.ResponseByContentType(context, http.StatusUnauthorized, utils.NewError(errors.New("unauthorized action")))
	}

	if err := handler.chatStore.DeleteMessage(cm); err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	messageResponse := newMessageResponse(context, cm, "message_deleted")
	curUser, _ := handler.GetCurrentUser(context)
	BroadcastMessage(chat, messageResponse, curUser.BlacklistedBy)

	return utils.ResponseByContentType(context, http.StatusOK, map[string]interface{}{"result": "ok"})
}

func (handler *Handler) UpdateMessage(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("message_id"), 10, 32)
	id := uint(id64)

	message, err := handler.chatStore.GetMessageByID(id)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	if message == nil {
		return utils.ResponseByContentType(context, http.StatusNotFound, utils.NotFound())
	}

	chat, err := handler.chatStore.GetById(message.ChatID)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}
	var cm model.Message

	req := &updateMessageRequest{}
	if err := req.bind(context, &cm); err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}

	cm.UserID = userIDFromToken(context)

	if cm.UserID == 0 {
		return utils.ResponseByContentType(context, http.StatusUnauthorized, utils.NewError(errors.New("unauthorized action")))
	}
	cm.ID = id

	if err = handler.chatStore.UpdateMessage(&cm); err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}
	messageResponse := newMessageResponse(context, &cm, "message_updated")
	curUser, _ := handler.GetCurrentUser(context)
	BroadcastMessage(chat, messageResponse, curUser.BlacklistedBy)

	return utils.ResponseByContentType(context, http.StatusCreated, messageResponse)
}

func (handler *Handler) ValidateParticipants(participants []uint) error  {
	if len(participants) < 1 {
		return fmt.Errorf("chat should have at least one participant")
	}

	if len(participants) > handler.config.MaxChatParticipants {
		return fmt.Errorf("chat room may have maximum %v participants", handler.config.MaxChatParticipants)
	}

	return nil
}

func (handler *Handler) GetCurrentUser(context echo.Context) (*model.User, error) {
	userID := userIDFromToken(context)
	return handler.userStore.GetByID(userID)
}
