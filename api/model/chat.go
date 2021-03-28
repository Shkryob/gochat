package model

import (
	"github.com/jinzhu/gorm"
)

type Chat struct {
	gorm.Model

	Admin      User
	AdminID    uint
	Title      string `gorm:"not null"`
	Users      []User `gorm:"many2many:chat_user;"`
	Messages   []Message
}

type Message struct {
	gorm.Model

	User   User
	UserID uint
	ChatID uint
	Body   string `gorm:"not null"`
}