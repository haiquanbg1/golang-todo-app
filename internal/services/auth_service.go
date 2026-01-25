package services

import (
	"fmt"

	"github.com/haiquanbg1/golang-todo-app/internal/models"
	"github.com/haiquanbg1/golang-todo-app/internal/repositories"
	"github.com/haiquanbg1/golang-todo-app/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username string, password string) (string, string, error)
	Register(username string, password string) error
}

type authService struct {
	userRepo repositories.UserRepository
	jwt      *utils.JWT
}

func NewAuthService(userRepo repositories.UserRepository, jwt *utils.JWT) AuthService {
	return &authService{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (service *authService) Login(username string, password string) (string, string, error) {
	user, err := service.userRepo.FindByUsername(username)
	if err != nil {
		return "", "", err
	}

	fmt.Println(user.ID)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", fmt.Errorf("Wrong password.")
	}

	accessToken, err := service.jwt.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", fmt.Errorf("Failed to generate access token")
	}

	refreshToken, err := service.jwt.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", fmt.Errorf("Failed to generate refresh token")
	}

	return accessToken, refreshToken, nil
}

func (service *authService) Register(username string, password string) error {
	user, _ := service.userRepo.FindByUsername(username)
	if user != nil {
		return fmt.Errorf("Username already exists.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user = &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return service.userRepo.Create(user)
}
