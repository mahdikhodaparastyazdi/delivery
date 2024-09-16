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
	Create(c context.Context, courior domain.COURIOR) (domain.COURIOR, error)
}
type queue3PL interface {
	Enqueue(msg dto.SendCourior, processAt *time.Time) error
}
type resolverProvier interface {
	ResolveCouriorProvider(providerName string) couriorproviders.CouriorSender
}
type Consumer struct {
	couriorRepo couriorRepository
	queue3PL    queue3PL
	resolver    resolverProvier
	logger      log.Logger
}

func New(logger log.Logger, q3pl queue3PL, resolver resolverProvier, couriorRepo couriorRepository) Consumer {
	return Consumer{
		logger:      logger,
		queue3PL:    q3pl,
		resolver:    resolver,
		couriorRepo: couriorRepo,
	}
}
