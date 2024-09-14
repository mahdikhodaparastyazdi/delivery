package middleware

import (
	"context"
	"delivery/internal/constants"
	"delivery/internal/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type clientService interface {
	GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error)
}

type InternalMiddleware struct {
	clientSvc clientService
}

func NewInternalMiddleware(cs clientService) InternalMiddleware {
	return InternalMiddleware{
		clientSvc: cs,
	}
}
func (am InternalMiddleware) respondForbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
}

func (am InternalMiddleware) Handle(c *gin.Context) {
	apiKey := c.GetHeader("ApiKey")

	if apiKey == "" {
		am.respondForbidden(c)
		return
	}
	client, err := am.clientSvc.GetClientByApiKey(c, apiKey)
	if err != nil {
		if errors.Is(err, constants.ErrWrongApiKey) {
			am.respondForbidden(c)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !client.IsActive {
		am.respondForbidden(c)
		return
	}

	c.Next()
}
