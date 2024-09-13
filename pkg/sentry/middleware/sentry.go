package middleware

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func SentryMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				sentry.CaptureException(err)
			}
		}
	}()
	c.Next()
	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			sentry.CaptureException(err.Err)
		}
	}
}
