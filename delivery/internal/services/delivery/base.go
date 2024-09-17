package delivery

import (
	recieved_task "delivery/internal/tasks/received_courior_status"
	send_task "delivery/internal/tasks/send_courior"
)

type Service struct {
	queue3PL  *send_task.Queue
	queueCore *recieved_task.Queue
}

func New(
	q3PL *send_task.Queue,
	qCore *recieved_task.Queue,
) Service {
	return Service{
		queue3PL:  q3PL,
		queueCore: qCore,
	}
}
