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

func GetUserByToken(db *gorm.DB, token string) (structs.User, error) {
	var user structs.User
	err := db.Where("token = ?", token).First(&user).Error
	return user, err
}

func GetUserByUsername(db *gorm.DB, username string) (structs.User, error) {
	var user structs.User
	err := db.Where("username = ?", username).First(&user).Error

	return user, err
}
