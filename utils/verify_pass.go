package utils

import "unicode"

func VerifyPass(pass string) bool {
	// Check password length
	if len(pass) < 3 {
		return false
	}

	var hasSymbol, hasNumber bool

	// Iterate through the password's characters
	for _, char := range pass {
		// Check if the character is a symbol
		if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			hasSymbol = true
		}

		// Check if the character is a number
		if unicode.IsDigit(char) {
			hasNumber = true
		}
	}

	// Return true if the password meets all criteria
	return hasSymbol && hasNumber
}
