package validator

func isValidEmail(email string) bool {
	return validateEmailRegex.MatchString(email)
}
