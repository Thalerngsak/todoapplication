package repository

import (
	"errors"
	"github.com/thalerngsak/todoapplication/model"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) UserRepository {
	return userStore{db: db}
}

func (u userStore) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := u.db.Where("id = ? ", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u userStore) GetByUserName(userName string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("username = ? ", userName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}
