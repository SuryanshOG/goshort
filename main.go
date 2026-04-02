package main

import (
	"goshort/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.HealthCheck)
	r.POST("/shorten", handler.ShortenURL)
	r.Run(":8080")
}
