package utils

import (
	"errors"
	"regexp"
)

var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	lower := regexp.MustCompile(`[a-z]`)
	if !lower.MatchString(password) {
		return errors.New("password must include at least one lowercase letter")
	}

	upper := regexp.MustCompile(`[A-Z]`)
	if !upper.MatchString(password) {
		return errors.New("password must include at least one uppercase letter")
	}

	number := regexp.MustCompile(`[0-9]`)
	if !number.MatchString(password) {
		return errors.New("password must include at least one number")
	}

	special := regexp.MustCompile(`[^A-Za-z0-9]`)
	if !special.MatchString(password) {
		return errors.New("password must include at least one special character")
	}

	return nil
}
