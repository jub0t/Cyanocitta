package main

import (
	"disco/api"
	database "disco/db"
	"disco/dfm"
	"disco/structs"
	// database "disco/db"
	// "disco/structs"
)

func main() {
	// Get the loaded envs & initialize database
	db := database.Connect()
	// Auto-migrate the databse to newer structure and start the REST API
	db.AutoMigrate(&structs.User{}, &structs.Bot{})
	dfm.PrepareSpace()

	api.Start(nil)
}
