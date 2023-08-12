package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	r := fiber.New(fiber.Config{
		AppName:       "Discochad",
		StrictRouting: true,
		CaseSensitive: true,
		Prefork:       true,
	})

	r.Get("/", func(ctx *fiber.Ctx) {
		ctx.(200, gin.H{
			"success": true,
		})
	})

	fmt.Println("Server Has Been Started")
	r.Listen(":8080")
}
