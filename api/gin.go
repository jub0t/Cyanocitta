package api

import (
	"disco/api/routes"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	r := fiber.New(fiber.Config{
		AppName:       "Discochad",
		StrictRouting: true,
		CaseSensitive: true,
		Prefork:       true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	r.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(ctx.App().Stack())
	})

	r.Post("/login", routes.LoginRoute(db))
	r.Post("/register", routes.RegisterRoute(db))
	r.Post("/create-bot", routes.CreateBotRoute(db))
	r.Post("/register", routes.DeleteBotRoute(db))

	fmt.Println("Server Has Been Started")
	r.Listen(":8080")
}
