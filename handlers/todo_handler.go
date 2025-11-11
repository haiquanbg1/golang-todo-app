package handlers

import (
	"net/http"

	"github.com/haiquanbg1/golang-todo-app/services"
)

type TodoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (handler *TodoHandler) Demo(w http.ResponseWriter, req *http.Request) {
	response := handler.todoService.Demo()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"` + response + `"}`))
}
