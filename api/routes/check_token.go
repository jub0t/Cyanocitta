package routes

import (
	"disco/structs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TokenMiddleware(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		headers := ctx.GetReqHeaders()
		token := headers["token"]

		if len(token) <= 0 {
			ctx.Status(403)
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Unauthenticated",
			})
		}

		return nil
	}
}
