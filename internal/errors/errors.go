package errors

import (
	"errors"
)

var (
	ErrRecordNotFound = errors.New("User not found")
	ErrForbidden      = errors.New("Access denied")
)
