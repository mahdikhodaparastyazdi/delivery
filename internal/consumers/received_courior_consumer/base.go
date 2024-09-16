package courior_consumer

import (
	"context"
	"delivery/internal/domain"
	"delivery/internal/dto"
	"time"

	log "delivery/pkg/logger"
)

type queueCore interface {
	Enqueue(msg dto.RecievedStatus, processAt *time.Time) error
}
type sendCouriorRepository interface {
	GetByID(c context.Context, id uint) (domain.COURIOR, error)
}

type Consumer struct {
	logger    log.Logger
	queueCore queueCore
}

func New(logger log.Logger, queue queueCore) Consumer {
	return Consumer{
		logger:    logger,
		queueCore: queue,
	}
}
