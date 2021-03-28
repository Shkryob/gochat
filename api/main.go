package main

import (
	"github.com/shkryob/gochat/config"
	"github.com/shkryob/gochat/db"
	"github.com/shkryob/gochat/handler"
	"github.com/shkryob/gochat/router"
	"github.com/shkryob/gochat/store"
)

func main() {
	r := router.New()

	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)
	configuration := config.ReadConfig()

	userStore := store.NewUserStore(d)
	chatStore := store.NewChatStore(d)
	h := handler.NewHandler(configuration, userStore, chatStore)
	h.Register(v1)
	r.Logger.Fatal(r.Start(":8081"))
}
