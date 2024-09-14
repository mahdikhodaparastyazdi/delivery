package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
	tasks "delivery/internal/tasks/send_courior"
)

type deliveryRepository interface {
	SendCourior(ctx context.Context, msg requests.SendCouriorRequest) error
	ReceiveCouriorStatus(ctx context.Context, msg requests.CouriorStatusRequest) error
}

type Service struct {
	deliveryRepository deliveryRepository
	queue3PL           *tasks.Queue
}

func New(
	sr deliveryRepository,
	q3PL *tasks.Queue,
) Service {
	return Service{
		deliveryRepository: sr,
		queue3PL:           q3PL,
	}
}
