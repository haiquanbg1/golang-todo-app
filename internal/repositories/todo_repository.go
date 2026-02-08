package repositories

import (
	"gorm.io/gorm"
)

type TodoRepository interface {
	Demo() string
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (repo *todoRepository) Demo() string {
	return "demo"
}
