package api

import (
	"net/http"

	"disco/config"
	database "disco/db"
	"disco/prom"
	"disco/structs"
	"disco/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var Config = config.GetConfig()

type CreateBotBody struct {
	Name     string
	Language int8
}

func CreateBotRoute(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		body := c.Get("Body").(structs.AnyData)
		user := c.Get("User").(structs.User)

		name, ok := body["Name"].(string)
		if !ok {
			return c.JSON(200, structs.Response{
				Success: false,
				Message: "Enter a valid name for the bot",
			})
		}

		if lang, ok := body["Language"].(int); !ok {
			println(lang)
			return c.JSON(200, structs.Response{
				Success: false,
				Message: "Choose a language for the bot",
			})
		} else {
			println(lang)
			if utils.IsLanguageValid(lang) {
				return c.JSON(200, structs.Response{
					Success: false,
					Message: "Invalid language for application",
				})
			}

			result := database.DB.Create(&structs.Bot{
				Name:        name,
				OwnerId:     uint(user.ID),
				AutoRestart: false,
				MaxRestarts: 0,
				BotId:       utils.RandomString(32),
				Language:    lang,
			})

			responseCode := http.StatusCreated
			message := "New bot created"
			if result.RowsAffected == 0 {
				responseCode = http.StatusOK
				message = "Error creating new bot"
			}

			return c.JSON(responseCode, structs.Response{
				Success: true,
				Message: message,
			})
		}
	}
}

func DeleteBotRoute(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		botId := c.Param("bot_id")

		user, ok := c.Request().Context().Value("User").(*structs.User)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, structs.Response{
				Success: false,
				Message: "User not found",
			})
		}

		owned := database.UserOwnsBot(db, uint(user.ID), botId)
		if !owned {
			// Bot Not Owned
			return echo.NewHTTPError(http.StatusUnauthorized, structs.Response{
				Success: false,
				Message: "Unauthorized Action",
			})
		}

		if len(botId) <= 0 {
			// Invalid BotId
			return c.JSON(http.StatusBadRequest, structs.Response{
				Success: false,
				Message: "Invalid BotId",
			})
		}

		err := database.DeleteBot(db, botId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, structs.Response{
				Success: false,
				Message: "Failed deleting bot",
			})
		}

		return c.NoContent(http.StatusOK)
	}
}

func StartBotRoute(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		botId := c.Param("bot_id")

		user, ok := c.Request().Context().Value("User").(*structs.User)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, structs.Response{
				Success: false,
				Message: "User not found",
			})
		}

		owned := database.UserOwnsBot(db, uint(user.ID), botId)
		if owned != true {
			// Bot Not Owned
			return echo.NewHTTPError(http.StatusUnauthorized, structs.Response{
				Success: false,
				Message: "Unauthorized Action",
			})
		}

		if len(botId) <= 0 {
			// Empty
			return c.JSON(http.StatusBadRequest, structs.Response{
				Success: false,
				Message: "Invalid BotId",
			})
		}

		bot, err := database.GetBotById(db, botId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, structs.Response{
				Success: false,
				Message: "Could not get Bot",
			})
		}

		processId := 0
		botPath := utils.PathJoin([]string{Config.StorePath, fmt.Sprintf("%v", botId)})
		startFile := utils.PathJoin([]string{botPath, bot.StartFile})

		switch lang := bot.Language; lang {
		case structs.JsLang:
			instance := structs.NodeInstance{
				IndexFile:     startFile,
				Name:          bot.Name,
				RestartOnStop: false,
				MaxRestarts:   0,
				CheckInterval: 5,
			}

			err = prom.StartInstance(instance)
			if err != nil {
				return c.JSON(http.StatusBadRequest, structs.Response{
					Success: false,
					Message: "Error occured while starting Node.js application",
					Data: structs.AnyData{
						"Error": err,
					},
				})
			}
		case structs.GoLang:
		case structs.PyLang:
		default:
			return c.JSON(http.StatusBadRequest, structs.Response{
				Success: false,
				Message: "Err! Please select a bot language.",
			})
		}

		return c.JSON(http.StatusOK, structs.ResponseAny{
			Success: true,
			Message: "Bot started successfully",
			Data: structs.AnyData{
				"StartFile": startFile,
				"Path":      botPath,
				"Bot":       bot,
				"ProcessId": processId,
			},
		})
	}
}
