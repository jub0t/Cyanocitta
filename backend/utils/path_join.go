package utils

import "path/filepath"

// Eg. ["/usr/bin/local", "../", "./local"] => /usr/bin/local
func PathJoin(components []string) string {
	return filepath.Clean(filepath.Join(components...))
}
