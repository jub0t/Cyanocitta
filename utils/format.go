package utils

import "fmt"

func FormatBytes(bytes int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	i := 0

	for bytes >= 1024 && i < len(sizes)-1 {
		bytes /= 1024
		i++
	}

	return fmt.Sprintf("%d %s", bytes, sizes[i])
}
