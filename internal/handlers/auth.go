package handlers

import "github.com/gin-gonic/gin"

func AuthHandler() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		name := ctx.PostForm("name")

		ctx.JSON(200, gin.H{
			"message": "Hello " + name,
		})
	})
}
