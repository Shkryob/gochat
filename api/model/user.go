package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"unique_index"`
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Image      *string
	Blacklists []Blacklist  `gorm:"foreignkey:FromID"`
	Friends    []Friend     `gorm:"foreignkey:FromID"`
	Chats      []Chat 		`gorm:"many2many:chat_user;"`

	BlacklistedBy []Blacklist `gorm:"foreignkey:ToID"`
}

type Blacklist struct {
	gorm.Model

	From    User
	FromID  uint `gorm:"primary_key" sql:"type:int not null"`
	To   	User
	ToID 	uint `gorm:"primary_key" sql:"type:int not null"`
}

type Friend struct {
	gorm.Model

	From    User
	FromID  uint `gorm:"primary_key" sql:"type:int not null"`
	To   	User
	ToID	uint `gorm:"primary_key" sql:"type:int not null"`
}
