package utils

import "regexp"

var (
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	hasLower      = regexp.MustCompile(`[a-z]`)
	hasUpper      = regexp.MustCompile(`[A-Z]`)
	hasDigit      = regexp.MustCompile(`\d`)
	hasSpecial    = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`)
)

func IsValidUsername(username string) bool {
	return usernameRegex.MatchString(username)
}

func IsValidPassword(password string) bool {
	return len(password) >= 8 &&
		hasLower.MatchString(password) &&
		hasUpper.MatchString(password) &&
		hasDigit.MatchString(password) &&
		hasSpecial.MatchString(password)
}
