package routes

import (
	database "disco/db"
	"disco/structs"
	"disco/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateBotBody struct {
	Name     string
	Language int8
}

func CreateBotRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Body CreateBotBody
		raw := ctx.Request().Body()

		if err := json.Unmarshal(raw, &Body); err != nil {
			// Error
			ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid JSON Body",
			})
		}

		// Retrieve the user information from the local context.
		User, ok := ctx.Locals("User").(structs.User)
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).
				SendString("User information not found in context")
		}

		if tx := db.Create(&structs.Bot{
			Name:        Body.Name,
			OwnerId:     User.ID,
			AutoRestart: false,
			MaxRestarts: 0,
			BotId:       utils.GenToken(32),
			Language:    Body.Language,
		}); tx.RowsAffected <= 0 {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Error creating new bot",
			})
		} else {
			return ctx.JSON(structs.Response{
				Success: true,
				Message: "New bot created",
			})
		}
	}
}

func DeleteBotRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(ctx.App().Stack())
	}
}

func StartBotRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bot_id := ctx.AllParams()["bot_id"]

		if len(bot_id) <= 0 {
			// Empty
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid BotId",
			})
		}

		if bot, err := database.GetBotById(db, bot_id); err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Could not get Bot",
			})
		} else {
			return ctx.JSON(structs.ResponseAny{
				Success: true,
				Data:    bot,
			})
		}
	}
}
