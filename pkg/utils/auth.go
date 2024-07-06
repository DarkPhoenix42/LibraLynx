package utils

import (
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

var bcrypt_cost int = 10

func HashPassword(password string) (string, error) {
	passwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt_cost)
	if err != nil {
		return "", err
	}
	return string(passwd), nil
}

func CheckUsername(username string) bool {

	username_regex := "^[a-zA-Z0-9_]{3,20}$"
	match, err := regexp.MatchString(username_regex, username)
	if err != nil {
		return false
	}
	return match
}

// https://stackoverflow.com/questions/25837241/password-validation-with-regexp
func CheckPassword(password string) bool {

	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 7 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func CheckEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
