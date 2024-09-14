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
	appEnv config.AppEnv
	logger log.Logger
}

func NewResolver(appEnv config.AppEnv, logger log.Logger) Resolver {
	return Resolver{
		appEnv: appEnv,
		logger: logger,
	}
}

func (r *Resolver) ResolveCouriorProvider(providerName string, templateID string) (couriorproviders.CouriorSender, error) {
	var (
		err    error
		driver couriorproviders.CouriorSender
	)

	switch providerName {
	case constants.COURIOR_PROVIDER1:
		driver = provider1.NewProvider1()
	case constants.COURIOR_PROVIDER2:
		driver = provider2.NewProvider2()
	default:
		err = constants.ErrProviderNotFound
	}

	return driver, err
}
