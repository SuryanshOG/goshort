package handler

import (
	"database/sql"
	"goshort/pkg/utils"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Laudasur :)",
	})
}

func ShortenURL(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	shortCode := utils.GenerateRandomShortCode(6)

	_, err := DB.Exec(
		"INSERT INTO urls(original_url, short_code) VALUES (?,?)",
		req.URL,
		shortCode,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to save"})
		return
	}

	c.JSON(200, gin.H{
		"short_url": "http://localhost:8080/" + shortCode,
	})
}
