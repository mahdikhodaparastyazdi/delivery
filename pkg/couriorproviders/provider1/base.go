package provider1

import (
	couriorproviders "delivery/pkg/couriorproviders"
	log "delivery/pkg/logger"
	"net/http"
	"time"
)

const Timeout = 60 * time.Second
const apiKeyHeader = "apiKey"

type provider1 struct {
	baseUrl3PL string
	apiKey3Pl  string
	hc         http.Client
	logger     log.Logger
}

func NewProvider1(baseUrl3PL, apiKey3Pl string, logger log.Logger) couriorproviders.CouriorSender {
	hc := http.Client{
		Timeout: Timeout,
	}
	return provider1{
		baseUrl3PL: baseUrl3PL,
		apiKey3Pl:  apiKey3Pl,
		hc:         hc,
		logger:     logger,
	}
}
