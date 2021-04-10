package handler

import (
	"github.com/labstack/gommon/log"
	"github.com/shkryob/gochat/utils"
	"os"
	"testing"

	"encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/shkryob/gochat/config"
	"github.com/shkryob/gochat/db"
	"github.com/shkryob/gochat/model"
	"github.com/shkryob/gochat/router"
	"github.com/shkryob/gochat/store"
)

var (
	d  *gorm.DB
	us UserStore
	cs ChatStore
	h  *Handler
	e  *echo.Echo
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func authHeader(token string) string {
	return "Token " + token
}

func setup() {
	d = db.TestDB()
	db.AutoMigrate(d)
	us = store.NewUserStore(d)
	cs = store.NewChatStore(d)
	conf := config.ReadConfig()
	h = NewHandler(conf, us, cs)
	h.config.MaxChatParticipants = 5
	e = router.New()
	loadFixtures()
}

func tearDown() {
	_ = d.Close()
	if err := db.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func responseMap(b []byte, key string) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m[key].(map[string]interface{})
}

func loadFixtures() error {
	u1 := model.User{
		Username: "user1",
		Email:    "user1@test.io",
	}
	u1.Password, _ = utils.HashPassword("secret")
	if err := us.Create(&u1); err != nil {
		return err
	}

	u2 := model.User{
		Username: "user2",
		Email:    "user2@test.io",
	}
	u2.Password, _ = utils.HashPassword("secret")
	if err := us.Create(&u2); err != nil {
		return err
	}

	a := model.Chat{
		Title:  "chat1 title",
		AdminID: 1,
	}
	cs.CreateChat(&a)
	cs.AddMessage(&a, &model.Message{
		Body:   "chat1 message1",
		UserID: 1,
	})
	users := []model.User{u1, u2}
	cs.ReplaceParticipants(&a, &users)

	a2 := model.Chat{
		Title:  "chat2 title",
		AdminID: 2,
	}
	cs.CreateChat(&a2)
	cs.AddMessage(&a2, &model.Message{
		Body:   "post2 comment1 by user1",
		UserID: 1,
	})
	cs.ReplaceParticipants(&a2, &users)

	return nil
}
