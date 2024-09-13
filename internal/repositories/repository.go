package repositories

import (
	"context"
	"delivery/internal/domain"
)

type CouriorRepository interface {
	Create(ctx context.Context, courior domain.COURIOR) (domain.COURIOR, error)
}

type ClientRepository interface {
	GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error)
}
