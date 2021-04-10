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
	h.config.MaxChatParticipants = 5
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
