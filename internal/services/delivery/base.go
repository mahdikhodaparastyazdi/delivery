package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
)

type deliveryRepository interface {
	SendCourior(ctx context.Context, msg requests.SendCouriorRequest) error
	ReceiveCouriorStatus(ctx context.Context, msg requests.CouriorStatusRequest) error
}

type Service struct {
	deliveryRepository deliveryRepository
}

func New(
	sr deliveryRepository,
) Service {
	return Service{
		deliveryRepository: sr,
	}
}
