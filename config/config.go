package config

import (
	"os"
)

type Config struct {
	StorePath string
}

func GetConfig() Config {
	return Config{
		StorePath: os.Getenv("STORE_PATH"),
	}
}
