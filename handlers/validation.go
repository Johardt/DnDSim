package handlers

import (
	"DnDSim/db"
	"regexp"
)

const MinPasswordLength = 3
const validDomains = `com|org|net|de|nl`

var emailRegex = regexp.MustCompile(`^[^@]+@[^@.]+\.(` + validDomains + `)$`)

type ValidationError struct {
	Field   string
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

func ValidateUsername(username string) error {
	if len(username) < 3 {
		return ValidationError{"username", "Username must be at least 3 characters long."}
	} else if user, _ := db.GetUserByName(username); user != nil {
		return ValidationError{"username", "Username already in use."}
	}
	return nil
}

func ValidateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return ValidationError{"email", "Invalid email address provided."}
	} else if user, _ := db.GetUserByEmail(email); user != nil {
		return ValidationError{"email", "Email address already in use."}
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < MinPasswordLength {
		return ValidationError{"password", "Password must be at least 3 characters long."}
	}
	return nil
}
