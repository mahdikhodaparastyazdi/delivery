package courior_consumer

import (
	"context"
	"delivery/internal/constants"
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
type couriorRepository interface {
	UpdateCouriorStatus(ctx context.Context,
		couriorId uint,
		status constants.CouriorStatus) error
}

type Consumer struct {
	logger            log.Logger
	queueCore         queueCore
	couriorRepository couriorRepository
	hc                http.Client
	baseUrlCore       string
	apiKeyCore        string
}

func New(logger log.Logger, queue queueCore,
	couriorRepo couriorRepository,
	coreBaseUrl,
	coreApiKey string) Consumer {
	hc := http.Client{
		Timeout: Timeout,
	}
	return Consumer{
		logger:            logger,
		queueCore:         queue,
		couriorRepository: couriorRepo,
		baseUrlCore:       coreBaseUrl,
		apiKeyCore:        coreApiKey,
		hc:                hc,
	}
}
