package services

import (
	"testing"

	"github.com/haiquanbg1/golang-todo-app/internal/models"
	"github.com/haiquanbg1/golang-todo-app/internal/repositories/mocks"
	"github.com/haiquanbg1/golang-todo-app/internal/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	jwtUtil := utils.NewJWT("test_secret_key", 15, 60)
	authService := NewAuthService(mockUserRepo, jwtUtil)

	username := "testuser"
	password := "testpassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockUserRepo.EXPECT().FindByUsername(username).Return(&models.User{
		ID:       1,
		Username: username,
		Password: string(hashedPassword),
	}, nil)

	accessToken, refreshToken, err := authService.Login(username, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)
}
