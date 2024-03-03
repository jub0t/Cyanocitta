package database

import (
	"disco/structs"

	"gorm.io/gorm"
)

func UserOwnsBot(db *gorm.DB, ownerId uint, botId string) bool {
	var count int64
	db.Model(&structs.Bot{}).Where("bot_id = ? AND owner_id = ?", botId, ownerId).Count(&count)

	return count > 0
}

func GetBotById(db *gorm.DB, id string) (structs.Bot, error) {
	var bot structs.Bot
	err := db.Where("bot_id = ?", id).First(&bot).Error

	return bot, err
}

func DeleteBot(db *gorm.DB, id string) error {
	var bot structs.Bot
	err := db.Where("bot_id = ?", id).First(&bot).Error

	return err
}
