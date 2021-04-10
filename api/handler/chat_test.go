package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/shkryob/gochat/router"
	"github.com/shkryob/gochat/router/middleware"
	"github.com/shkryob/gochat/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListChatsCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	e := router.New()
	req := httptest.NewRequest(echo.GET, "/api/chats", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	assert.NoError(t, jwtMiddleware(func(context echo.Context) error {
		return h.GetChats(c)
	})(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var aa chatListResponse
		err := json.Unmarshal(rec.Body.Bytes(), &aa)
		assert.NoError(t, err)
		assert.Equal(t, 2, aa.ChatsCount)
	}
}

func TestCreateChatCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"chat":{"title":"chat3", "participants":[1,2]}}`
	)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.POST, "/api/chats", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := jwtMiddleware(func(context echo.Context) error {
		return h.CreateChat(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		var a SingleChatResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "chat3", a.Chat.Title)
		assert.Equal(t, "user1", a.Chat.Admin.Username)
		assert.Equal(t, 2, len(a.Chat.Participants))
	}
}

func TestUpdateChatCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"chat":{"title":"chat3new", "participants":[1,2]}}`
	)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.PUT, "/api/chats/:chat_id", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/chats/:chat_id")
	c.SetParamNames("chat_id")
	c.SetParamValues("1")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.UpdateChat(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var a SingleChatResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "chat3new", a.Chat.Title)
	}
}

func TestDeleteChatCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.DELETE, "/api/chats/:chat_id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/chats/:chat_id")
	c.SetParamNames("chat_id")
	c.SetParamValues("1")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.DeleteChat(c)
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetMessagesCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.GET, "/api/chats/:chat_id/messages", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(2)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/chats/:chat_id/messages")
	c.SetParamNames("chat_id")
	c.SetParamValues("1")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.GetMessages(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var cc messageListResponse
		err := json.Unmarshal(rec.Body.Bytes(), &cc)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(cc.Messages))
	}
}

func TestAddMessageCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"message":{"body":"chat1 message2 by user2"}}`
	)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.POST, "/api/chats/:chat_id/messages", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(2)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/chats/:chat_id/messages")
	c.SetParamNames("chat_id")
	c.SetParamValues("1")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.AddMessage(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		var c SingleMessageResponse
		err := json.Unmarshal(rec.Body.Bytes(), &c)
		assert.NoError(t, err)
		assert.Equal(t, "chat1 message2 by user2", c.Message.Body)
		assert.Equal(t, "user2", c.Message.User.Username)
	}
}

func TestDeleteMessageCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.DELETE, "/api/chats/:chat_id/messages/:message_id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/chats/:chat_id/messages/:message_id")
	c.SetParamNames("chat_id")
	c.SetParamValues("1")
	c.SetParamNames("message_id")
	c.SetParamValues("1")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.DeleteMessage(c)
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
