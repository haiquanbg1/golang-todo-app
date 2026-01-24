package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/haiquanbg1/golang-todo-app/internal/constants"
	"github.com/haiquanbg1/golang-todo-app/internal/errors"
	"github.com/haiquanbg1/golang-todo-app/internal/models"
)

type TodoRepository interface {
	FindByUser(status constants.TodoStatus, userId uint, limit int, offset int) ([]*models.Todo, error)
	FindById(id uint) (*models.Todo, error)
	Create(todo *models.Todo, userId uint) error
	Update(id uint, updates constants.TodoInput) error
	Delete(id uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (repo *todoRepository) FindByUser(status constants.TodoStatus, userId uint, limit int, offset int) ([]*models.Todo, error) {
	var todos []*models.Todo

	fmt.Println(userId)
	query := repo.db.Where("user_id = ?", userId)
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Order("created_at desc").Offset(offset).Limit(limit).Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (repo *todoRepository) FindById(id uint) (*models.Todo, error) {
	var todo *models.Todo

	err := repo.db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}

	return todo, nil
}

func (repo *todoRepository) Create(todo *models.Todo, userId uint) error {
	todo.UserID = userId
	return repo.db.Create(todo).Error
}

func (repo *todoRepository) Update(id uint, updates constants.TodoInput) error {
	var todo *models.Todo

	err := repo.db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrRecordNotFound
		}
		return err
	}

	return repo.db.Model(&todo).Updates(updates).Error
}

func (repo *todoRepository) Delete(id uint) error {
	var todo *models.Todo

	err := repo.db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrRecordNotFound
		}
		return err
	}

	return repo.db.Delete(&todo).Error
}
