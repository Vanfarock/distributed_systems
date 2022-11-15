package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	count := 0
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, count)
	})

	r.GET("/sum", func(ctx *gin.Context) {
		count += 1
	})

	r.Run(":8080")
}
