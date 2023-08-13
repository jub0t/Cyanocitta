package utils

import "regexp"

func VerifyPass(pass string) bool {
	// Atleast Three Chars, One Symbol & One Number
	pattern := `^(?=.*[!@#$%^&*()_+{}\[\]:;<>,.?~])(?=.*[0-9])(.{3,})$`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return regex.MatchString(pass)
}
