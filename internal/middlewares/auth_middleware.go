package middlewares

import (
	"context"
	"net/http"

	"github.com/haiquanbg1/golang-todo-app/internal/repositories"
	"github.com/haiquanbg1/golang-todo-app/internal/utils"
)

type AuthMiddleware struct {
	jwt            *utils.JWT
	userRepository repositories.UserRepository
}

func NewAuthMiddleware(jwt *utils.JWT, userRepository repositories.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		jwt:            jwt,
		userRepository: userRepository,
	}
}

func (authMiddleware *AuthMiddleware) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessCookie, err := r.Cookie("accessToken")
			if err == nil {
				userId, err := authMiddleware.jwt.ParseAccessToken(accessCookie.Value)
				if err == nil {
					user, err := authMiddleware.userRepository.FindById(userId)
					if err != nil {
						http.Error(w, "Unauthorized", http.StatusUnauthorized)
						return
					}

					ctx := context.WithValue(r.Context(), "user", user)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}

			refreshCookie, err := r.Cookie("refreshToken")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userId, err := authMiddleware.jwt.ParseRefreshToken(refreshCookie.Value)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			newAccessToken, err := authMiddleware.jwt.GenerateAccessToken(userId)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "accessToken",
				Value:    newAccessToken,
				HttpOnly: true,
				MaxAge:   15 * 60,
				Path:     "/",
			})

			user, err := authMiddleware.userRepository.FindById(userId)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
