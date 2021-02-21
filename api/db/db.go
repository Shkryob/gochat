package db

import (
	"github.com/labstack/gommon/log"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/shkryob/gochat/model"
	_ "gorm.io/driver/mysql"
)

func New() *gorm.DB {
	db, err := gorm.Open("mysql", "chat:secret@tcp(db:3306)/chat")
	if err != nil {
		log.Fatal("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gochat_test.db")
	if err != nil {
		log.Fatal("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./gochat_test.db"); err != nil {
		return err
	}
	return nil
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Chat{},
		&model.Message{},
	)
}
