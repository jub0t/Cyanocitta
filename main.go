package main

import (
	"disco/structs"
	"disco/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	gin.SetMode(gin.ReleaseMode)

	dsn := utils.MakeMysqlDsn(utils.DsnConfig{
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("DBUSER"),
		Pass:     os.Getenv("DBPASS"),
		Host:     os.Getenv("DBHOST"),
		Database: os.Getenv("DBNAME"),
	})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&structs.User{})

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
