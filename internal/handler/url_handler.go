package handler

import "github.com/gin-gonic/gin"

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is running :)",
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

	shortCode := "abc123"

	c.JSON(200, gin.H{
		"short_url": "http://localhost:8080/" + shortCode,
	})
}
