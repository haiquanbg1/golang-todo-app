package errors

import (
	"errors"
)

var (
	ErrRecordNotFound = errors.New("Not found")
	ErrForbidden      = errors.New("Access denied")
	ErrConflict       = errors.New("Conflict")
	ErrUnauthorized   = errors.New("Unauthorized")
	ErrInvalidInput   = errors.New("Invalid input")
)
