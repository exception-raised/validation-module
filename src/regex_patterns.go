package validator

import "regexp"

const (
	validateEmailPattern          string = "^[\\w\\-.+]+@([\\w-]+\\.)+[\\w-]{2,}$"
	validateStrictPasswordPattern string = "^[a-zA-Z0-9!@#$%^&*()_+{}\\[\\]:;<>,.?~\\-]*[A-Z][a-zA-Z0-9!@#$%^&*()_+{}\\[\\]:;<>,.?~\\-]*[0-9][a-zA-Z0-9!@#$%^&*()_+{}\\[\\]:;<>,.?~\\-]*[!@#$%^&*()_+{}\\[\\]:;<>,.?~\\-][a-zA-Z0-9!@#$%^&*()_+{}\\[\\]:;<>,.?~\\-]*$"
)

func compileRegex(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

var (
	validateEmailRegex          = compileRegex(validateEmailPattern)
	validateStrictPasswordRegex = compileRegex(validateEmailPattern)
)
