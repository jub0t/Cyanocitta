package main

import (
	"disco/api"
	"disco/config"
	// database "disco/db"
	// "disco/structs"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	// Get the loaded envs & initialize database
	config := config.GetConfig()
	// db := database.GetDB()

	// Auto-migrate the databse to newer structure and start the REST API
	// db.AutoMigrate(&structs.User{}, &structs.Bot{})
	api.Start(nil, &config)
}
