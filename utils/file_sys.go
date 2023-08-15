package utils

import "os"

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true // File exists
	}

	if os.IsNotExist(err) {
		return false // File does not exist
	}

	return false // Error occurred during file stat
}
