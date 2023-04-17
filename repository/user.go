package repository

import (
	"github.com/thalerngsak/todoapplication/model"
)

type UserRepository interface {
	GetByID(id uint) (*model.User, error)
	GetByUserName(userName string) (*model.User, error)
}
