package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	r.Run(":8080")
}