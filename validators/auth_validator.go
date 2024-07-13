package validators

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

func AuthValidator(email, password string) error {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	matchEmail, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	if err != nil {
		log.Printf("Error matching email regex: %v", err)
		return errors.New("error validating email")
	}
	if !matchEmail {
		return errors.New("invalid email")
	}

	// Validate password length
	if len(password) < 8 {
		return errors.New("password must have at least 8 characters")
	}

	// Validate password for at least one uppercase letter
	matchUppercase, err := regexp.MatchString(`[A-Z]`, password)
	if err != nil || !matchUppercase {
		return errors.New("password must have at least one uppercase letter")
	}

	// Validate password for at least one lowercase letter
	matchLowercase, err := regexp.MatchString(`[a-z]`, password)
	if err != nil || !matchLowercase {
		return errors.New("password must have at least one lowercase letter")
	}

	// Validate password for at least one digit
	matchDigit, err := regexp.MatchString(`\d`, password)
	if err != nil || !matchDigit {
		return errors.New("password must have at least one number")
	}

	return nil
}
