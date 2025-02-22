package validator

func IsValidEmail(email string) bool {
	return isValidEmail(email)
}

func IsValidPassword(password string, minLength int, maxLength int, regex string) bool {
	return isValidPassword(password, minLength, maxLength, regex)
}
