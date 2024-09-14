package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
)

func (s Service) SendCourior(ctx context.Context, msg requests.SendCouriorRequest) error {
	return s.deliveryRepository.SendCourior(ctx, msg)
}
