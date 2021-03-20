package store

import (
	"github.com/jinzhu/gorm"
	"github.com/shkryob/gochat/model"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetByID(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(u *model.User) (err error) {
	return us.db.Create(u).Error
}

func (us *UserStore) GetByEmail(e string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) GetByUsername(username string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Username: username}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Update(u *model.User) error {
	return us.db.Model(u).Update(u).Error
}

func (us *UserStore) List(offset, limit int, search string) ([]model.User, int, error) {
	var (
		users []model.User
		count int
	)

	us.db.Model(&users).Where("username LIKE ?", "%" + search + "%").Count(&count)
	us.db.Offset(offset).
		Limit(limit).
		Where("username LIKE ?", "%" + search + "%").
		Order("created_at desc").Find(&users)

	return users, count, nil
}

func (us *UserStore) GetBlacklist(fromID uint, toID uint) (*model.Blacklist, error) {
	var m model.Blacklist
	if err := us.db.Where(&model.Blacklist{FromID: fromID, ToID: toID}).
		First(&m).Error; err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) GetFriend(fromID uint, toID uint) (*model.Friend, error) {
	var m model.Friend
	if err := us.db.Where(&model.Friend{FromID: fromID, ToID: toID}).
		First(&m).Error; err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) CreateBlackList(bl *model.Blacklist) (err error) {
	return us.db.Create(bl).Error
}

func (us *UserStore) RemoveBlackList(bl *model.Blacklist) (err error) {
	return us.db.Delete(bl).Error
}

func (us *UserStore) CreateFriend(fr *model.Friend) (err error) {
	return us.db.Create(fr).Error
}

func (us *UserStore) RemoveFriend(fr *model.Friend) (err error) {
	return us.db.Delete(fr).Error
}
