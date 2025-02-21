package validator

import "regexp"

const (
	validateEmailPattern string = "^[\\w\\-.+]+@([\\w-]+\\.)+[\\w-]{2,}$"
)

func compileRegex(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

var (
	validateEmailRegex = compileRegex(validateEmailPattern)
)
