package main

import (
	"goshort/internal/handler"

	"github.com/gin-gonic/gin"

	"goshort/internal/db"
)

func main() {
	r := gin.Default()
	conn := db.ConnectDB()
	handler.DB = conn
	r.GET("/", handler.HealthCheck)
	r.POST("/shorten", handler.ShortenURL)
	r.Run(":8080")
}
