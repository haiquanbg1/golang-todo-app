package repositories

import (
	"fmt"

	"github.com/haiquanbg1/golang-todo-app/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id int) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) FindById(id int) (*models.User, error) {
	var user models.User

	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User not found.")
		}

		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User not found.")
		}

		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}
