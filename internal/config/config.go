package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	JWT_SECRET           string
	PORT                 string
	ENV                  string
	DSN                  string
	ACCESS_TOKEN_EXPIRY  int
	REFRESH_TOKEN_EXPIRY int
}

func Load() AppConfig {
	_ = godotenv.Load(".env")

	secret := getEnv("JWT_SECRET", "todo-app-secret")

	port := getEnv("PORT", "8080")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	environment := getEnv("ENVIRONMENT", "development")

	default_dsn := "root@tcp(127.0.0.1:3306)/todo_app?parseTime=true&loc=Local"
	dsn := getEnv("DATABASE_DSN", default_dsn)

	accessTokenExpiry := getEnvInt("ACCESS_TOKEN_EXPIRY", 900)
	refreshTokenExpiry := getEnvInt("REFRESH_TOKEN_EXPIRY", 604800)

	return AppConfig{
		JWT_SECRET:           secret,
		PORT:                 port,
		ENV:                  environment,
		DSN:                  dsn,
		ACCESS_TOKEN_EXPIRY:  accessTokenExpiry,
		REFRESH_TOKEN_EXPIRY: refreshTokenExpiry,
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && strings.TrimSpace(v) != "" {
		return v
	}
	return def
}

func getEnvInt(key string, def int) int {
	if v, ok := os.LookupEnv(key); ok {
		var i int
		if _, err := fmt.Sscanf(v, "%d", &i); err == nil && i >= 0 {
			return i
		}
	}
	return def
}
