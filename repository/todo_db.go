package repository

import (
	"errors"
	"github.com/thalerngsak/todoapplication/model"
	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")

type todoRepository struct {
	db *gorm.DB
}

func NewTodoDB(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) Update(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(model.Todo{}).Error
}

func (r *todoRepository) GetByID(id uint, userID uint) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) List(userID uint) ([]model.Todo, error) {
	var todos []model.Todo
	if err := r.db.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return todos, nil
}
