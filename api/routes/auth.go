package routes

import (
	database "disco/db"
	"disco/structs"
	"disco/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthBody struct {
	Username string
	Password string
}

func LoginRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Body AuthBody
		raw := ctx.Request().Body()

		if err := json.Unmarshal(raw, &Body); err != nil {
			// Error
			ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid JSON Body",
			})
		}

		username := Body.Username
		if len(username) < 3 {
			ctx.JSON(structs.Response{
				Success: false,
				Message: "Username must be atleast 3 characters",
			})
		}

		if user, err := database.GetUserByUsername(db, Body.Username); err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Unable to fetch user info",
			})
		} else {
			if (user != structs.User{}) {
				return ctx.JSON(structs.Response{
					Success: false,
					Data: structs.AnyData{
						"Token": user.Token,
					},
				})
			} else {
				return ctx.JSON(structs.Response{
					Success: false,
					Message: "Unable to fetch user info",
				})
			}
		}

		return nil
	}
}

func RegisterRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Body AuthBody
		raw := ctx.Request().Body()

		if err := json.Unmarshal(raw, &Body); err != nil {
			// Error
			ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid JSON Body",
			})
		}

		username := Body.Username
		if len(username) < 3 {
			ctx.JSON(structs.Response{
				Success: false,
				Message: "Username must be atleast 3 characters",
			})
		}

		if database.UserExists(db, username) {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Username is taken",
			})
		}

		password := Body.Password

		if !utils.VerifyPass(password) {
			// Password Invalid
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Password must match criteria",
			})
		}

		token := utils.GenToken(128)
		db.Create(&structs.User{
			Username: username,
			Password: password,
			Token:    token,
		})

		return ctx.JSON(structs.Response{
			Success: true,
			Message: "Registered new account",
			Data: structs.AnyData{
				"Token": token,
			},
		})
	}
}
