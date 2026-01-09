package utils

import "regexp"

var (
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	passwordRegex = regexp.MustCompile(`^.{6,}$`)
)

func IsValidUsername(username string) bool {
	return usernameRegex.MatchString(username)
}

func IsValidPassword(password string) bool {
	return passwordRegex.MatchString(password)
}
