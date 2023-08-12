package api

import "github.com/gin-gonic/gin"

func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.PureJSON(200, gin.H{
			"success": true,
		})
	})

	r.Run()
}
