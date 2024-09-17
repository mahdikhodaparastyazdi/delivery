package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/webhook", func(c *gin.Context) {
		c.String(http.StatusOK, "nil")
	})

	router.Run(":8080")
}
