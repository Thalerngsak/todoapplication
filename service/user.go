package service

import (
	"github.com/thalerngsak/todoapplication/model"
)

type UserService interface {
	GetByID(userID uint) (*model.User, error)
	GetByUserName(userName string) (*model.User, error)
}
