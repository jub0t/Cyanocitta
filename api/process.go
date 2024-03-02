package api


import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
  "disco/prom"
  "fmt"
)

func ProcessResourcesRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
    pid := ctx.Params("pid")
    status, err := prom.GetProcessStatus(pid)
    
    if err != nil {
      fmt.Printf("%s", err)
      return ctx.JSON("error broski, could not find process")
    }

		return ctx.JSON(status)
	}
}
