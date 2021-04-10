package handler

import (
	"github.com/shkryob/gochat/config"
	"github.com/shkryob/gochat/model"
)

type Handler struct {
	config    config.Configuration
	userStore UserStore
	chatStore ChatStore
}

type ChatStore interface {
	GetById(uint) (*model.Chat, error)
	List(int, int, *model.User) ([]model.Chat, int, error)
	CreateChat(*model.Chat) error
	UpdateChat(*model.Chat) error
	DeleteChat(*model.Chat) error
	ReplaceParticipants(*model.Chat, *[]model.User) error

	AddMessage(*model.Chat, *model.Message) error
	GetMessagesByChatId(uint, *model.User) ([]model.Message, error)
	GetMessageByID(uint) (*model.Message, error)
	DeleteMessage(*model.Message) error
	UpdateMessage(*model.Message) error
}

type UserStore interface {
	GetByID(uint) (*model.User, error)
	GetByIDs([]uint) []model.User
	GetByEmail(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	List(int, int, string) ([]model.User, int, error)
	GetBlacklist(fromID uint, toID uint) (*model.Blacklist, error)
	GetFriend(fromID uint, toID uint) (*model.Friend, error)
    CreateBlackList(bl *model.Blacklist) (err error)
	RemoveBlackList(bl *model.Blacklist) (err error)
	CreateFriend(bl *model.Friend) (err error)
	RemoveFriend(bl *model.Friend) (err error)
}

func NewHandler(config config.Configuration, us UserStore, cs ChatStore) *Handler {
	return &Handler{
		config:    config,
		userStore: us,
		chatStore: cs,
	}
}
