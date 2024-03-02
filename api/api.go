package api

import (
	"disco/config"
	"disco/structs"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	StartTime int = time.Now().Second()
	Port      int = 8080
)

func Start(db *gorm.DB, conf *config.Config) {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(structs.Response{
			Success: true,
			Data: structs.AnyData{
				"Uptime": time.Now().Second() - StartTime,
			},
		})
	})

	// Middlewares
	r.Use(Serdeware)

	// Authentication
	r.POST("/login", LoginRoute(db))
	r.POST("/register", RegisterRoute(db))

	// Bot Manage
	r.POST("/create-bot", TokenMiddleware(db), CreateBotRoute(db))
	r.POST("/start-bot/:bot_id", TokenMiddleware(db), StartBotRoute(db, conf))
	r.POST("/delete-bot/:bot_id", TokenMiddleware(db), DeleteBotRoute(db))
	r.POST("/process-resources/:pid", ProcessResourcesRoute(db))

	fmt.Printf("Server Should Be Available http://localhost:%v\n", Port)
	r.Start(fmt.Sprintf("127.0.0.1:%v", Port))
}
