package database

import (
	"disco/structs"

	"gorm.io/gorm"
)

func UserExists(db *gorm.DB, username string) bool {
	var userCount int64
	db.Model(&structs.User{}).Where("username = ?", username).Count(&userCount)

	return userCount > 0
}
