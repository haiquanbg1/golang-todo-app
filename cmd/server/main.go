package main

import (
	"log"
	"net/http"

	"github.com/haiquanbg1/golang-todo-app/internal/config"
	handlers "github.com/haiquanbg1/golang-todo-app/internal/handlers/rest"
	"github.com/haiquanbg1/golang-todo-app/internal/repositories"
	"github.com/haiquanbg1/golang-todo-app/internal/services"
	"github.com/haiquanbg1/golang-todo-app/internal/utils"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// load config and connect, migrate database
	cfg := config.Load()
	_, err := utils.Connect(cfg.DSN)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// config router
	router := chi.NewRouter()
	router.Use(chimw.Logger)
	router.Use(chimw.Recoverer)

	// setup dependencies
	todoRepository := repositories.NewTodoRepository()
	todoService := services.NewTodoService(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoService)

	// setup routes
	router.Get("/demo", todoHandler.Demo)

	log.Printf("listening on %s", cfg.PORT)
	log.Fatal(http.ListenAndServe(cfg.PORT, router))
}
