package webFramework

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// install gin
// go get -u github.com/gin-gonic/gin
func TestGin() {
	fmt.Println("----------------------------------------Start Gin----------------------------------------")
	// Create a Gin router with default middleware: logger and recovery (crash-free) middleware
	router := gin.Default()

	// Define a route for GET requests to the root URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Define a route for GET requests to /ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Define a route for POST requests to /submit
	router.POST("/submit", func(c *gin.Context) {
		var json struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required"`
		}

		// Bind JSON payload to the struct and check for errors
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Respond with the received data
		c.JSON(http.StatusOK, gin.H{
			"name":  json.Name,
			"email": json.Email,
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
	fmt.Println("----------------------------------------End Gin  ----------------------------------------")
}
