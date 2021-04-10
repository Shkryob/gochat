package store

import (
	"github.com/jinzhu/gorm"
	"github.com/shkryob/gochat/model"
)

type ChatStore struct {
	db *gorm.DB
}

func NewChatStore(db *gorm.DB) *ChatStore {
	return &ChatStore{
		db: db,
	}
}

func (chatStore *ChatStore) GetById(id uint) (*model.Chat, error) {
	var m model.Chat

	err := chatStore.db.Where(id).Preload("Users").Preload("Admin").First(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, err
	}

	return &m, err
}

func (chatStore *ChatStore) List(offset, limit int, user *model.User) ([]model.Chat, int, error) {
	var (
		chats []model.Chat
		count int
	)

	chatStore.db.Model(&chats).
		Joins("inner join chat_user on chat_user.chat_id = chats.id").
		Where("chat_user.user_id > ?", user.ID, user.ID).
		Count(&count)

	chatStore.db.Offset(offset).
		Limit(limit).
		Preload("Users").
		Preload("Admin").
		Order("created_at desc").
		Joins("inner join chat_user on chat_user.chat_id = chats.id").
		Where("chat_user.user_id > ?", user.ID, user.ID).
		Find(&chats)

	return chats, count, nil
}

func (chatStore *ChatStore) CreateChat(a *model.Chat) error {
	tx := chatStore.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(a.ID).Preload("Admin").Preload("Users").Find(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (chatStore *ChatStore) UpdateChat(a *model.Chat) error {
	tx := chatStore.db.Begin()
	if err := tx.Model(a).Update(a).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(a.ID).Preload("Admin").Preload("Users").Find(a).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (chatStore *ChatStore) ReplaceParticipants(a *model.Chat, participants *[]model.User) error {
	tx := chatStore.db.Begin()
	if err := tx.Model(a).Update(a).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Model(a).Association("Users").Replace(participants)

	if err := tx.Where(a.ID).Preload("Admin").Preload("Users").Find(a).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (chatStore *ChatStore) DeleteChat(a *model.Chat) error {
	return chatStore.db.Delete(a).Error
}

func (chatStore *ChatStore) AddMessage(a *model.Chat, c *model.Message) error {
	err := chatStore.db.Model(a).Association("Messages").Append(c).Error
	if err != nil {
		return err
	}

	return chatStore.db.Where(c.ID).Preload("User").First(c).Error
}

func (chatStore *ChatStore) UpdateMessage(m *model.Message) error {
	tx := chatStore.db.Begin()
	if err := tx.Model(m).Update(m).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(m.ID).Preload("User").Find(m).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (chatStore *ChatStore) GetMessagesByChatId(id uint, curUser *model.User) ([]model.Message, error) {
	var c model.Chat
	blacklist := curUser.Blacklists
	var blacklistIDs []uint
	for _, user := range blacklist {
		blacklistIDs = append(blacklistIDs, user.ToID)
	}

	query := chatStore.db.Where(id)
	if len(blacklist) > 0 {
		query = query.Preload(
			"Messages", "user_id NOT IN (?)", blacklistIDs,
		)
	} else {
		query = query.Preload("Messages")
	}
	err := query.Preload("Messages.User").First(&c).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, err
	}

	return c.Messages, nil
}

func (chatStore *ChatStore) GetMessageByID(id uint) (*model.Message, error) {
	var m model.Message
	if err := chatStore.db.Where(id).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, err
	}

	return &m, nil
}

func (chatStore *ChatStore) DeleteMessage(c *model.Message) error {
	return chatStore.db.Delete(c).Error
}
