package service

import (
	"github.com/thalerngsak/todoapplication/model"
	"github.com/thalerngsak/todoapplication/repository"
	"time"
)

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) Create(userID uint, message string) error {
	todo := &model.Todo{
		Message:   message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Done:      false,
		UserID:    userID,
	}

	return s.repo.Create(todo)
}

func (s *todoService) Update(id uint, userID uint, message string, done bool) error {
	todo, err := s.repo.GetByID(id, userID)
	if err != nil {
		return err
	}

	todo.Message = message
	todo.UpdatedAt = time.Now()
	todo.Done = done

	return s.repo.Update(todo)
}

func (s *todoService) Delete(id uint, userID uint) error {
	return s.repo.Delete(id, userID)
}

func (s *todoService) GetByID(id uint, userID uint) (*model.Todo, error) {
	return s.repo.GetByID(id, userID)
}

func (s *todoService) List(userID uint) ([]model.Todo, error) {
	return s.repo.List(userID)
}
