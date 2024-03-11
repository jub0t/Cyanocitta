package dfm

import (
	"fmt"
	"os"
	"path/filepath"
)

func Exists(dirPath string) bool {
	_, err := os.Stat(filepath.Clean(dirPath)) // filepath.Clean removes redundant separators

	switch {
	case err == nil:
		return true
	case os.IsNotExist(err):
		return false
	default:
		fmt.Printf("Unexpected error occurred checking dir '%s': %v\n", dirPath, err)
		return false
	}
}
