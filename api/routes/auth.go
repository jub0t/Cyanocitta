package routes

import (
	"disco/structs"
	"disco/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoginRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(ctx.App().Stack())
	}
}

type RegisterBody struct {
	Username string
	Password string
}

func RegisterRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Body RegisterBody
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

		password := Body.Password
		pass_valid := utils.VerifyPass(password)

		if !pass_valid {
			// Password Invalid
			ctx.JSON(structs.Response{
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
			Data: structs.AnyData{
				"Token": token,
			},
		})
	}
}
