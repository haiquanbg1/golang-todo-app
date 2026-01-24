package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/haiquanbg1/golang-todo-app/internal/constants"
	"github.com/haiquanbg1/golang-todo-app/internal/errors"
	"github.com/haiquanbg1/golang-todo-app/internal/models"
	"github.com/haiquanbg1/golang-todo-app/internal/services"
	"github.com/haiquanbg1/golang-todo-app/internal/utils"
)

type TodoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (handler *TodoHandler) FindByUser(w http.ResponseWriter, req *http.Request) {
	user, ok := req.Context().Value("user").(*models.User)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	status := req.URL.Query().Get("status")
	page := req.URL.Query().Get("page")
	limit := req.URL.Query().Get("limit")

	todoPage, todoLimit := utils.ParsePaginationParams(page, limit)
	todoOffset := (todoPage - 1) * todoLimit

	todos, err := handler.todoService.FindByUser(status, user.ID, todoLimit, todoOffset)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to fetch todos")
		return
	}

	utils.WriteResponse(w, http.StatusOK, todos)
}

func (handler *TodoHandler) FindById(w http.ResponseWriter, req *http.Request) {
	user, ok := req.Context().Value("user").(*models.User)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idParam := chi.URLParam(req, "id")
	if idParam == "" {
		utils.WriteError(w, http.StatusBadRequest, "Missing todo ID")
		return
	}

	todoId, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	todo, err := handler.todoService.FindById(uint(todoId), user.ID)
	if err != nil {
		switch err {
		case errors.ErrForbidden:
			utils.WriteError(w, http.StatusForbidden, "Access denied")
			return
		case errors.ErrRecordNotFound:
			utils.WriteError(w, http.StatusNotFound, "Todo not found")
			return
		default:
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.WriteResponse(w, http.StatusOK, todo)
}

func (handler *TodoHandler) Create(w http.ResponseWriter, req *http.Request) {
	user, ok := req.Context().Value("user").(*models.User)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var input *constants.TodoInput
	if err := utils.ParseJSON(req, &input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if input.Task == nil {
		utils.WriteError(w, http.StatusBadRequest, "Task is required")
		return
	}

	todo := &models.Todo{
		Task:        *input.Task,
		Description: *input.Description,
		Status:      constants.TodoStatus(*input.Status),
	}

	if err := handler.todoService.Create(todo, user.ID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create todo")
		return
	}

	utils.WriteResponse(w, http.StatusCreated, todo)
}

func (handler *TodoHandler) Update(w http.ResponseWriter, req *http.Request) {
	user, ok := req.Context().Value("user").(*models.User)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idParam := chi.URLParam(req, "id")
	if idParam == "" {
		utils.WriteError(w, http.StatusBadRequest, "Missing todo ID")
		return
	}

	todoId, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	var input constants.TodoInput
	if err := utils.ParseJSON(req, &input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := handler.todoService.Update(uint(todoId), input, user.ID); err != nil {
		switch err {
		case errors.ErrForbidden:
			utils.WriteError(w, http.StatusForbidden, "Access denied")
			return
		default:
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.WriteResponse(w, http.StatusNoContent, nil)
}

func (handler *TodoHandler) Delete(w http.ResponseWriter, req *http.Request) {
	user, ok := req.Context().Value("user").(*models.User)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idParam := chi.URLParam(req, "id")
	if idParam == "" {
		utils.WriteError(w, http.StatusBadRequest, "Missing todo ID")
		return
	}

	todoId, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	if err := handler.todoService.Delete(uint(todoId), user.ID); err != nil {
		switch err {
		case errors.ErrForbidden:
			utils.WriteError(w, http.StatusForbidden, "Access denied")
			return
		default:
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.WriteResponse(w, http.StatusNoContent, nil)
}
