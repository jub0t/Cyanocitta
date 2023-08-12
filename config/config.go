package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	StorePath string
}

func GetConfig() Config {
	godotenv.Load()

	return Config{
		StorePath: os.Getenv("STORE_PATH"),
	}
}
