package database

import (
	"disco/structs"
	"disco/utils"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB() *gorm.DB {
	dsn := utils.MakeMysqlDsn(utils.DsnConfig{
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("DBUSER"),
		Pass:     os.Getenv("DBPASS"),
		Host:     os.Getenv("DBHOST"),
		Database: os.Getenv("DBNAME"),
	})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
	}

	// Debug Only
	db.Delete(&structs.User{})

	return db
}
