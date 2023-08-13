package utils

import (
	"regexp"
)

func VerifyPass(pass string) bool {
	// At least Three Chars, One Symbol & One Number
	pattern := `^(?=.*[A-Za-z])(?=.*[0-9])(?=.*[^A-Za-z0-9]).{3,}$`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return regex.MatchString(pass)
}
