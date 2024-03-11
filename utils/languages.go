package utils

import "disco/structs"

func IsLanguageValid(lang int) bool {
	for _, value := range structs.Langauges {
		if value == lang {
			return true
		}
	}

	return false
}
