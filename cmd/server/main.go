package main

import (
	"log"
	"net/http"

	"github.com/haiquanbg1/golang-todo-app/internal/config"
	handlers "github.com/haiquanbg1/golang-todo-app/internal/handlers/rest"
	"github.com/haiquanbg1/golang-todo-app/internal/middlewares"
	"github.com/haiquanbg1/golang-todo-app/internal/repositories"
	"github.com/haiquanbg1/golang-todo-app/internal/services"
	"github.com/haiquanbg1/golang-todo-app/internal/utils"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// load config and connect, migrate database
	cfg := config.Load()

	jwt := utils.NewJWT(cfg.JWT_SECRET, cfg.ACCESS_TOKEN_EXPIRY, cfg.REFRESH_TOKEN_EXPIRY)

	db, err := utils.Connect(cfg.DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// config router
	router := chi.NewRouter()
	router.Use(chimw.Logger)
	router.Use(chimw.Recoverer)

	// setup dependencies
	todoRepository := repositories.NewTodoRepository(db)
	userRepository := repositories.NewUserRepository(db)

	todoService := services.NewTodoService(todoRepository)
	authService := services.NewAuthService(userRepository, jwt)

	todoHandler := handlers.NewTodoHandler(todoService)
	authHandler := handlers.NewAuthHandler(authService, cfg.ACCESS_TOKEN_EXPIRY, cfg.REFRESH_TOKEN_EXPIRY)

	authMiddleware := middlewares.NewAuthMiddleware(jwt, userRepository)

	// setup routes
	router.Post("/api/v1/login", authHandler.Login)
	router.Post("/api/v1/register", authHandler.Register)

	router.Route("/api/v1/todos", func(r chi.Router) {
		r.Use(authMiddleware.Middleware())

		r.Get("/", todoHandler.FindByUser)
		r.Get("/{id}", todoHandler.FindById)
		r.Post("/", todoHandler.Create)
		r.Patch("/{id}", todoHandler.Update)
		r.Delete("/{id}", todoHandler.Delete)
	})

	log.Printf("listening on %s", cfg.PORT)
	log.Fatal(http.ListenAndServe(cfg.PORT, router))
}
