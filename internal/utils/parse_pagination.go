package utils

import (
	"strconv"

	"github.com/haiquanbg1/golang-todo-app/internal/constants"
)

func ParsePaginationParams(pageStr, limitStr string) (int, int) {
	page := constants.DefaultPage
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	limit := constants.DefaultPageSize
	if l, err := strconv.Atoi(limitStr); err == nil {
		if l > constants.MaxPageSize {
			limit = constants.MaxPageSize
		} else if l > 0 {
			limit = l
		}
	}

	return page, limit
}
