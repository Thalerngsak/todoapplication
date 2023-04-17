package repository

import (
	"github.com/thalerngsak/todoapplication/model"
)

type TodoRepository interface {
	Create(todo *model.Todo) error
	Update(todo *model.Todo) error
	Delete(id uint, userID uint) error
	GetByID(id uint, userID uint) (*model.Todo, error)
	List(userID uint) ([]model.Todo, error)
}
