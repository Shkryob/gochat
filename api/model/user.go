package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
