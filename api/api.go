package api

import (
	"disco/config"
	database "disco/db"
	"disco/structs"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	StartTime int64 = time.Now().UnixMilli()
	Port      int   = 8080
)

func Start(conf *config.Config) {
	db := database.DB
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200, structs.Response{
			Success: true,
			Message: "Success",
		})
	})

	// Middlewares
	r.Use(Serdeware)
	r.Use(CustomBodyParser)

	// Authentication
	r.POST("/login", LoginRoute(db))
	r.POST("/register", RegisterRoute(db))

	// Bot Manage
	r.POST("/create-bot", CreateBotRoute(db), TokenMiddleware)
	r.POST("/start-bot/:bot_id", StartBotRoute(db), TokenMiddleware)
	r.POST("/delete-bot/:bot_id", DeleteBotRoute(db), TokenMiddleware)
	r.GET("/process-resources/:pid", ProcessResourcesRoute())

	fmt.Printf("Server Should Be Available http://localhost:%v\n", Port)
	r.Start(fmt.Sprintf("127.0.0.1:%v", Port))
}
