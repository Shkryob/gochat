package handler

import (
	"errors"
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

	return utils.ResponseByContentType(context, http.StatusOK, newChatResponse(context, chat))
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

	err := handler.chatStore.CreateChat(&chat)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return utils.ResponseByContentType(context, http.StatusCreated, newChatResponse(context, &chat))
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

	return utils.ResponseByContentType(context, http.StatusOK, newChatResponse(context, chat))
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

	return utils.ResponseByContentType(context, http.StatusCreated, newMessageResponse(context, &cm))
}

func (handler *Handler) GetMessages(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("chat_id"), 10, 32)
	id := uint(id64)

	cm, err := handler.chatStore.GetMessagesByChatId(id)
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

	return utils.ResponseByContentType(context, http.StatusOK, map[string]interface{}{"result": "ok"})
}

func (handler *Handler) UpdateMessage(context echo.Context) error {
	id64, err := strconv.ParseUint(context.Param("message_id"), 10, 32)
	id := uint(id64)

	chat, err := handler.chatStore.GetById(id)
	if err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	if chat == nil {
		return utils.ResponseByContentType(context, http.StatusNotFound, utils.NotFound())
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

	if err = handler.chatStore.UpdateMessage(&cm); err != nil {
		return utils.ResponseByContentType(context, http.StatusInternalServerError, utils.NewError(err))
	}

	return utils.ResponseByContentType(context, http.StatusCreated, newMessageResponse(context, &cm))
}
