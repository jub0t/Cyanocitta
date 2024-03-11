package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	StorePath string
}

func GetConfig() Config {
	// Load environment variables from .env file
	godotenv.Load()

	return Config{
		StorePath: os.Getenv("STORE_PATH"),
	}
}

var C = GetConfig()
