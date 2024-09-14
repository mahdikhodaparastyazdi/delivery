package middleware

import (
	"net/http"

	log "delivery/pkg/logger"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func SentryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					sentry.CaptureException(err)
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()
		c.Next()
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if err.IsType(gin.ErrorTypePublic) {
					log.Info("Public error", log.J{
						"error": err.Err.Error(),
					})
					continue
				}
				sentry.CaptureException(err.Err)
			}
		}
	}
}
