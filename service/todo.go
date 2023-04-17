package service

import (
	"github.com/thalerngsak/todoapplication/model"
)

type TodoService interface {
	Create(userID uint, message string) error
	Update(id uint, userID uint, message string, done bool) error
	Delete(id uint, userID uint) error
	GetByID(id uint, userID uint) (*model.Todo, error)
	List(userID uint) ([]model.Todo, error)
}
