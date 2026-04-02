package handler

import (
	"database/sql"
	"goshort/pkg/utils"
	"net/http"

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

	if req.URL == "" {
		c.JSON(400, gin.H{"error": "url is required"})
		return
	}

	shortCode := utils.GenerateRandomShortCode(6)

	_, err := DB.Exec(
		"INSERT INTO urls(original_url, short_code) VALUES (?, ?)",
		req.URL,
		shortCode,
	)

	if err != nil {
		var existingCode string
		err2 := DB.QueryRow(
			"SELECT short_code FROM urls WHERE original_url = ?",
			req.URL,
		).Scan(&existingCode)

		if err2 == nil {
			c.JSON(200, gin.H{
				"short_url": "http://localhost:8080/" + existingCode,
			})
			return
		}

		c.JSON(500, gin.H{"error": "failed to save"})
		return
	}

	c.JSON(200, gin.H{
		"short_url": "http://localhost:8080/" + shortCode,
	})
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	var originalURL string

	err := DB.QueryRow(
		"SELECT original_url FROM urls WHERE short_code = ?",
		code,
	).Scan(&originalURL)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if originalURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid URL"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
