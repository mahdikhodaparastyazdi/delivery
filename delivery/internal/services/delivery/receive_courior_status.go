package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
	"delivery/internal/dto"
)

func (s Service) ReceiveCouriorStatus(ctx context.Context, msg requests.CouriorStatusRequest) error {
	var dtoMessage = dto.RecievedStatus{
		DeliveryID: msg.DeliverId,
		Status:     msg.Status,
	}
	return s.queueCore.Enqueue(dtoMessage, nil)
}
