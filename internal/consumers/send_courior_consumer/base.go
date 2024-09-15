package courior_consumer

import (
	"context"
	"delivery/internal/domain"
	"delivery/internal/dto"
	"time"

	"delivery/pkg/couriorproviders"
	log "delivery/pkg/logger"
)

type couriorRepository interface {
	GetByID(c context.Context, id uint) (domain.COURIOR, error)
}
type queue3PL interface {
	Enqueue(msg dto.SendCourior, processAt *time.Time) error
}
type resolverProvier interface {
	ResolveCouriorProvider(providerName string) (couriorproviders.CouriorSender, error)
}
type Consumer struct {
	queue3PL queue3PL
	resolver resolverProvier
	logger   log.Logger
}

func New(logger log.Logger, q3pl queue3PL, resolver resolverProvier) Consumer {
	return Consumer{
		logger:   logger,
		queue3PL: q3pl,
		resolver: resolver,
	}
}
