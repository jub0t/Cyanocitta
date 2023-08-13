package api

import (
	"disco/api/routes"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	r := fiber.New(fiber.Config{
		StrictRouting:         true,
		CaseSensitive:         true,
		Prefork:               true,
		DisableDefaultDate:    true,
		DisableStartupMessage: true,
		AppName:               "Discochad",
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})

	r.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(ctx.App().Stack())
	})

	r.Post("/login", routes.LoginRoute(db))
	r.Post("/register", routes.RegisterRoute(db))
	r.Post("/create-bot", routes.CreateBotRoute(db))
	r.Post("/delete-bot/:bot_id", routes.DeleteBotRoute(db))

	r.Listen(":8080")
}
