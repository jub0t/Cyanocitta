package database

import (
	"disco/structs"
	"disco/utils"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := utils.MakeMysqlDsn(utils.DsnConfig{
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("DBUSER"),
		Pass:     os.Getenv("DBPASS"),
		Host:     os.Getenv("DBHOST"),
		Database: os.Getenv("DBNAME"),
	})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// Debug Only
	db.Delete(&structs.User{})

	return db
}
