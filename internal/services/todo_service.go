package services

import (
	"github.com/haiquanbg1/golang-todo-app/internal/repositories"
)

type TodoService interface {
	Demo() string
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{
		repo: repo,
	}
}

func (service *todoService) Demo() string {
	return service.repo.Demo()
}
