package services

import (
	"github.com/haiquanbg1/golang-todo-app/internal/constants"
	"github.com/haiquanbg1/golang-todo-app/internal/errors"
	"github.com/haiquanbg1/golang-todo-app/internal/models"
	"github.com/haiquanbg1/golang-todo-app/internal/repositories"
)

type TodoService interface {
	FindByUser(status string, userId uint, limit int, offset int) ([]*models.Todo, error)
	FindById(id uint, userId uint) (*models.Todo, error)
	Create(todo *models.Todo, userId uint) error
	Update(id uint, updates constants.TodoInput, userId uint) error
	Delete(id uint, userId uint) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{
		repo: repo,
	}
}

func (service *todoService) FindByUser(status string, userId uint, limit int, offset int) ([]*models.Todo, error) {
	return service.repo.FindByUser(constants.TodoStatus(status), userId, limit, offset)
}

func (service *todoService) FindById(id uint, userId uint) (*models.Todo, error) {
	todo, err := service.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if todo.UserID != userId {
		return nil, errors.ErrForbidden
	}

	return service.repo.FindById(id)
}

func (service *todoService) Create(todo *models.Todo, userId uint) error {
	return service.repo.Create(todo, userId)
}

func (service *todoService) Update(id uint, updates constants.TodoInput, userId uint) error {
	todo, err := service.repo.FindById(id)
	if err != nil {
		return err
	}

	if todo.UserID != userId {
		return errors.ErrForbidden
	}

	return service.repo.Update(id, updates)
}

func (service *todoService) Delete(id uint, userId uint) error {
	todo, err := service.repo.FindById(id)
	if err != nil {
		return err
	}

	if todo.UserID != userId {
		return errors.ErrForbidden
	}

	return service.repo.Delete(id)
}
