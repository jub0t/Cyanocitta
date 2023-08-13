package main

import (
	"disco/api"
	"disco/config"
	database "disco/db"
	"disco/structs"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config := config.GetConfig()
	fmt.Println(config)

	db := database.GetDbConnection()
	db.AutoMigrate(&structs.User{})
	api.Start(db)
}
