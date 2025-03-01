package validator

import (
	"log"
	"regexp"
)

// improvements:
// provide error message for invalid inputs

func isValidEmail(email string) bool {
	return validateEmailRegex.MatchString(email)
}

func isValidLength(entry string, minLength int, maxLength int) bool {
	if len(entry) < minLength || len(entry) > maxLength {
		return false
	}
	return true
}

func isValidPassword(password string, minLength int, maxLength int, regex string) bool {
	if !isValidLength(password, minLength, maxLength) {
		return false
	}

	if regex == "" {
		return true
	}

	compiledRegex, err := regexp.Compile(regex)
	if err != nil {
		log.Fatalf("error compiling regex: %v", err)
		return false
	}
	return compiledRegex.MatchString(password)
}

func isNonEmptyString(str string) bool {
	return len(str) > 0
}
