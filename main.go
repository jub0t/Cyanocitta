package main

import (
	"disco/api"
	database "disco/db"
	"disco/structs"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := database.GetDbConnection()
	db.AutoMigrate(&structs.User{})
	api.Start(db)
}
