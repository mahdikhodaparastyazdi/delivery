package courior_resolver

import (
	"delivery/internal/config"
	"delivery/internal/constants"
	"delivery/pkg/couriorproviders"
	provider1 "delivery/pkg/couriorproviders/provider1"
	provider2 "delivery/pkg/couriorproviders/provider2"
	log "delivery/pkg/logger"
)

type Resolver struct {
	cfg    config.Config
	logger log.Logger
}

func NewResolver(cfg config.Config, logger log.Logger) Resolver {
	return Resolver{
		cfg:    cfg,
		logger: logger,
	}
}

func (r Resolver) ResolveCouriorProvider(providerName string) couriorproviders.CouriorSender {
	var driver couriorproviders.CouriorSender
	switch providerName {
	case constants.COURIOR_PROVIDER1:
		driver = provider1.NewProvider1(r.cfg.CouriorBaseUrl, r.cfg.APIKey3PL, r.logger)
	case constants.COURIOR_PROVIDER2:
		driver = provider2.NewProvider2()
	default:
		driver = provider1.NewProvider1(r.cfg.CouriorBaseUrl, r.cfg.APIKey3PL, r.logger)
	}

	return driver
}
