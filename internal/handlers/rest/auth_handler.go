package rest

import (
	"fmt"
	"net/http"

	"github.com/haiquanbg1/golang-todo-app/internal/services"
	"github.com/haiquanbg1/golang-todo-app/internal/utils"
)

type AuthHandler struct {
	authService        services.AuthService
	accessTokenExpiry  int
	refreshTokenExpiry int
}

func NewAuthHandler(authService services.AuthService, accessTokenExpiry int, refreshTokenExpiry int) *AuthHandler {
	return &AuthHandler{
		authService:        authService,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

func (handler *AuthHandler) Login(w http.ResponseWriter, req *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := utils.ParseJSON(req, &input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validateCredentials(input.Username, input.Password); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := handler.authService.Login(input.Username, input.Password)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		MaxAge:   handler.accessTokenExpiry,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		MaxAge:   handler.refreshTokenExpiry,
		HttpOnly: true,
	})

	utils.WriteResponse(w, http.StatusOK, map[string]string{"message": "Login successful"})
}

func (handler *AuthHandler) Register(w http.ResponseWriter, req *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := utils.ParseJSON(req, &input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validateCredentials(input.Username, input.Password); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := handler.authService.Register(input.Username, input.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, map[string]string{"message": "Registration successful. Please log in."})
}

func validateCredentials(username, password string) error {
	if username == "" {
		return fmt.Errorf("Username is required.")
	}

	if password == "" {
		return fmt.Errorf("Password is required.")
	}

	if !utils.IsValidUsername(username) {
		return fmt.Errorf("Invalid username.")
	}

	if !utils.IsValidPassword(password) {
		return fmt.Errorf("Invalid password.")
	}

	return nil
}
