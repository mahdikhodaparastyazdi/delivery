package courior_consumer

import (
	"context"
	"delivery/internal/domain"

	log "delivery/pkg/logger"
)

type couriorRepository interface {
	GetByID(c context.Context, id uint) (domain.COURIOR, error)
}

type Consumer struct {
	logger log.Logger
}

func New(logger log.Logger) Consumer {
	return Consumer{
		logger: logger,
	}
}
