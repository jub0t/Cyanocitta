package api

import (
	"disco/api/routes"
	"disco/structs"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var StartTime int64 = time.Now().UnixNano()

func Start(db *gorm.DB) {
	r := fiber.New(fiber.Config{
		// Settings For Speed
		StrictRouting:         true,
		CaseSensitive:         true,
		Prefork:               false, // Don't Enable
		DisableDefaultDate:    true,
		DisableStartupMessage: true,
		AppName:               "Discochad",

		// Faster JSON
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	r.Get("/", func(ctx *fiber.Ctx) error {
		now := time.Now().UnixNano()
		start := int64(StartTime)

		return ctx.JSON(structs.Response{
			Success: true,
			Message: "Server Is Up & Working!",
			Data: structs.AnyData{
				"UptimeMS": (now - start) / 1000 / 1000,
				"UptimeNS": now - start,
			},
		})
	})

	r.Post("/login", routes.LoginRoute(db))
	r.Post("/register", routes.RegisterRoute(db))

	// Bot Manage
	r.Post("/create-bot", routes.TokenMiddleware(db), routes.CreateBotRoute(db))
	r.Post("/start-bot/:bot_id", routes.TokenMiddleware(db), routes.StartBotRoute(db))
	r.Post("/delete-bot/:bot_id", routes.TokenMiddleware(db), routes.DeleteBotRoute(db))

	r.Listen(":8080")
}
