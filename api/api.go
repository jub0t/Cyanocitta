package api

import (
	"disco/config"
	"disco/structs"
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	StartTime int64 = time.Now().UnixNano()
	Port      int   = 8080
)

func Start(db *gorm.DB, conf *config.Config) {
	r := fiber.New(fiber.Config{
		// Settings For Speed
		AppName:               "Cyanocitta",
		StrictRouting:         true,
		CaseSensitive:         true,
		Prefork:               false, // Don't Enable
		DisableDefaultDate:    true,
		DisableStartupMessage: true,

		// Faster JSON
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	r.Get("/", func(ctx *fiber.Ctx) error {
		now := time.Now().UnixNano()

		return ctx.JSON(structs.Response{
			Success: true,
			Data: structs.AnyData{
				"Uptime": now - StartTime,
			},
		})
	})

	// Authentication
	r.Post("/login", LoginRoute(db))
	r.Post("/register", RegisterRoute(db))

	// Bot Manage
	r.Post("/create-bot", TokenMiddleware(db), CreateBotRoute(db))
	r.Post("/start-bot/:bot_id", TokenMiddleware(db), StartBotRoute(db, conf))
	r.Post("/delete-bot/:bot_id", TokenMiddleware(db), DeleteBotRoute(db))
  r.Get("/process-resources/:pid", ProcessResourcesRoute(db))

	fmt.Printf("Server Should Be Available http://localhost:%v\n", Port)
	r.Listen(fmt.Sprintf("127.0.0.1:%v", Port))
}
