package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateBotRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("User")

		return ctx.JSON(user)
	}
}

func DeleteBotRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(ctx.App().Stack())
	}
}

func StartBotRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(ctx.App().Stack())
	}
}
