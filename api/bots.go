package api

import (
	"disco/config"
	database "disco/db"
	"disco/prom"
	"disco/structs"
	"disco/utils"
	"encoding/json"
	"fmt"

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
			return ctx.Status(fiber.StatusInternalServerError).JSON(structs.Response{
				Success: false,
				Message: "User not found",
			})
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

func StartBotRoute(db *gorm.DB, conf *config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bot_id := ctx.AllParams()["bot_id"]

		User, ok := ctx.Locals("User").(structs.User)
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).JSON(structs.Response{
				Success: false,
				Message: "User not found",
			})
		}

		if owned := database.UserOwnsBot(db, User.ID, bot_id); owned != true {
			// Bot Not Owned
			return ctx.Status(fiber.StatusUnauthorized).JSON(structs.Response{
				Success: false,
				Message: "Unauthorized Action",
			})
		}

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
			var process_id int64
			bot_path := utils.PathJoin([]string{conf.StorePath, fmt.Sprintf("/%v", bot_id)})
			start_file := utils.PathJoin([]string{bot_path, bot.StartFile})

			switch lang := bot.Language; lang {
			case structs.JsLang:
				{
					// Javascript
					instance := structs.NodeInstance{
						IndexFile:     start_file,
						Name:          bot.Name,
						RestartOnStop: false,
						MaxRestarts:   0,
						CheckInterval: 5,
					}

					prom.StartInstance(instance)
					fmt.Println(instance)
				}
			case structs.GoLang:
				{
				}
			case structs.Pylang:
				{
				}
			default:
				{
					return ctx.JSON(structs.Response{
						Success: false,
						Message: "Err! Please select a bot language.",
					})
				}
			}

			return ctx.JSON(structs.ResponseAny{
				Success: true,
				Message: "Bot Has Been Started!",
				Data: structs.AnyData{
					"StartFile": start_file,
					"Path":      bot_path,
					"Bot":       bot,
					"ProcessId": process_id,
				},
			})
		}
	}
}
