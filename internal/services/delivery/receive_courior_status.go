package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
)

func (s Service) ReceiveCouriorStatus(ctx context.Context, msg requests.CouriorStatusRequest) error {
	return s.deliveryRepository.ReceiveCouriorStatus(ctx, msg)
}
