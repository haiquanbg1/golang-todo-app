package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	JWT_SECRET string
	PORT       string
	ENV        string
}

func Load() AppConfig {
	_ = godotenv.Load(".env")

	secret := getEnv("JWT_SECRET", "todo-app-secret")

	port := getEnv("PORT", "8080")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	environment := getEnv("ENVIRONMENT", "development")

	return AppConfig{
		JWT_SECRET: secret,
		PORT:       port,
		ENV:        environment,
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && strings.TrimSpace(v) != "" {
		return v
	}
	return def
}
