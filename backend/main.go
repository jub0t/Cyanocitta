package main

import (
	"disco/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.PureJSON(200, gin.H{
			"success": true,
		})
	})

	store_path := os.Getenv("STORE_PATH")
	test_path := utils.PathJoin([]string{store_path, "./test/"})
	fmt.Println(test_path)

	r.Run()
}
