package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
	"github.com/shkryob/gochat/model"
	"github.com/shkryob/gochat/router/middleware"
	"github.com/shkryob/gochat/utils"
	"sync"

	"github.com/labstack/echo/v4"
)

var connectionPool = struct {
	sync.RWMutex
	connections map[uint][]*websocket.Conn
}{
	connections: make(map[uint][]*websocket.Conn),
}

var (
	upgrader = websocket.Upgrader{}
)

func (handler *Handler) GetSocket(context echo.Context) error {
	ws, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		context.Logger().Error(err)
		return err
	}
	defer ws.Close()

	for {
		if err != nil {
			context.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			context.Logger().Error(err)
			return err
		}
		msgData := utils.JsonToMap(msg)
		if msgData["action"] == "authorize" {
			authorizeSocket(fmt.Sprintf("%v", msgData["jwt"]), ws)
		}
	}
}

func authorizeSocket(jwtToken string, ws *websocket.Conn)  {
	config := middleware.JWTConfig{}
	config.SigningKey= utils.JWTSecret

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.SigningKey, nil
	})

	if err != nil {
		log.Info(err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["id"].(float64))
		val, ok := connectionPool.connections[userID]
		if ok {
			connectionPool.connections[userID] = append(val, ws)
		} else {
			var val []*websocket.Conn
			connectionPool.connections[userID] =  append(val, ws)
		}
	}
}

func BroadcastMessage(chat *model.Chat, message *SingleMessageResponse)  {
	for _, user := range chat.Users {
		if sockets, ok := connectionPool.connections[user.ID]; ok {
			for _, socket := range sockets {
				b, _ := json.MarshalIndent(message, "", "\t")
				if err := socket.WriteMessage(websocket.TextMessage, b); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}