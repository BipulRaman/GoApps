package controllers

import (
	"net/http"

	"go-gin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

// GetTest handles GET requests to /test
func (t TestController) GetTest(c *gin.Context) {
	// Return a simple JSON Item
	item := models.Item{
		ID:   1,
		Name: "Test Item",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GET test endpoint hit",
		"item":    item,
	})
}

// CreateTest handles POST requests to /test
func (t TestController) CreateTest(c *gin.Context) {
	var requestBody models.Item
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body, expected Item type",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "POST test endpoint hit",
		"data":    requestBody,
	})
}
