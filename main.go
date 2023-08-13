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

	db := database.GetDB()
	db.AutoMigrate(&structs.User{}, &structs.Bot{})
	api.Start(db)
}
