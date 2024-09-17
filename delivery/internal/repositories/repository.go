package repositories

import (
	"context"
	"delivery/internal/constants"
	"delivery/internal/domain"
)

type DeliveryRepository interface {
	UpdateCouriorStatus(ctx context.Context,
		couriorId uint,
		status constants.CouriorStatus) error
	Create(c context.Context, courior domain.COURIOR) (domain.COURIOR, error)
}

type ClientRepository interface {
	GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error)
}
