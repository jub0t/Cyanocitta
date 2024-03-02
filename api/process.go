package api

import (
	"disco/prom"
	"disco/structs"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProcessResourcesRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		pid := ctx.Params("pid")
		status, err := prom.GetProcessStatus(pid)

		if err != nil {
			fmt.Printf("%s", err)
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Process not found, could not fetch resource data",
			})
		}

		return ctx.JSON(structs.ResponseAny{
			Data:    status,
			Success: true,
		})
	}
}
