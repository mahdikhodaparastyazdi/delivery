package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
	recieved_task "delivery/internal/tasks/received_courior_status"
	send_task "delivery/internal/tasks/send_courior"
)

type deliveryRepository interface {
	SendCourior(ctx context.Context, msg requests.SendCouriorRequest) error
	ReceiveCouriorStatus(ctx context.Context, msg requests.CouriorStatusRequest) error
}

type Service struct {
	deliveryRepository deliveryRepository
	queue3PL           *send_task.Queue
	queueCore          *recieved_task.Queue
}

func New(
	sr deliveryRepository,
	q3PL *send_task.Queue,
	qCore *recieved_task.Queue,
) Service {
	return Service{
		deliveryRepository: sr,
		queue3PL:           q3PL,
		queueCore:          qCore,
	}
}
