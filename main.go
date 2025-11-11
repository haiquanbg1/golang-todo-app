package main

import (
	"log"
	"net/http"

	"github.com/haiquanbg1/golang-todo-app/config"
	"github.com/haiquanbg1/golang-todo-app/handlers"
	"github.com/haiquanbg1/golang-todo-app/repositories"
	"github.com/haiquanbg1/golang-todo-app/services"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// load config
	cfg := config.Load()

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
