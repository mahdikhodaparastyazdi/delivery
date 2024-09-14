package repositories

import (
	"context"
	"delivery/internal/api/rest/requests"
	"delivery/internal/domain"
)

type DeliveryRepository interface {
	Create(ctx context.Context, courior domain.COURIOR) (domain.COURIOR, error)
	SendCourior(ctx context.Context, msg requests.SendCouriorRequest) error
	ReceiveCouriorStatus(ctx context.Context, msg requests.CouriorStatusRequest) error
}

type ClientRepository interface {
	GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error)
}
