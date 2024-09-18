package courior_consumer

import (
	"context"
	"delivery/internal/constants"
	"delivery/internal/domain"
	"delivery/internal/dto"
	"net/http"
	"time"

	log "delivery/pkg/logger"
)

const Timeout = 60 * time.Second
const apiKeyHeader = "apiKey"

type queueCore interface {
	Enqueue(msg dto.RecievedStatus, processAt *time.Time) error
}
type queue3PL interface {
	Enqueue(msg dto.SendCourior, processAt *time.Time) error
}
type couriorRepository interface {
	UpdateCouriorStatus(ctx context.Context,
		deliveryID uint,
		status constants.CouriorStatus) error
	GetById(ctx context.Context,
		deliveryID uint,
	) (domain.Delivery, error)
}

type Consumer struct {
	logger            log.Logger
	queueCore         queueCore
	queue3PL          queue3PL
	couriorRepository couriorRepository
	hc                http.Client
	baseUrlCore       string
	apiKeyCore        string
}

func New(logger log.Logger,
	queueCore queueCore,
	queue3PL queue3PL,
	couriorRepo couriorRepository,
	coreBaseUrl,
	coreApiKey string) Consumer {
	hc := http.Client{
		Timeout: Timeout,
	}
	return Consumer{
		logger:            logger,
		queueCore:         queueCore,
		queue3PL:          queue3PL,
		couriorRepository: couriorRepo,
		baseUrlCore:       coreBaseUrl,
		apiKeyCore:        coreApiKey,
		hc:                hc,
	}
}
