package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/webhook", func(c *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(100)
		// If the random number is less than 5, return an error (5% chance)
		if randomNum < 5 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "service not available",
			})
			return
		}
		c.String(http.StatusOK, "nil")
	})

	router.Run(":8080")
}
