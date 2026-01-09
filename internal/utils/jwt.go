package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret             []byte
	accessTokenExpiry  int
	refreshTokenExpiry int
}

func NewJWT(secret string, accessTokenExpiry int, refreshTokenExpiry int) *JWT {
	return &JWT{
		secret:             []byte(secret),
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

func (j *JWT) GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    "access",
		"exp":     time.Now().Add(time.Duration(j.accessTokenExpiry) * time.Second).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWT) GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    "refresh",
		"exp":     time.Now().Add(time.Duration(j.refreshTokenExpiry) * time.Second).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWT) Parse(tokenStr string, expectType string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return j.secret, nil
	})

	if err != nil {
		return 0, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims")
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != expectType {
		return 0, fmt.Errorf("token type mismatch: expected %s", expectType)
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found or invalid")
	}

	return uint(userIDFloat), nil
}

func (j *JWT) ParseAccessToken(tokenStr string) (uint, error) {
	return j.Parse(tokenStr, "access")
}

func (j *JWT) ParseRefreshToken(tokenStr string) (uint, error) {
	return j.Parse(tokenStr, "refresh")
}
