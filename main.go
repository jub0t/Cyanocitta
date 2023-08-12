package main

import (
	database "disco/db"
	"disco/structs"
)

func main() {
	db := database.GetDbConnection()
	db.AutoMigrate(&structs.User{})
}
