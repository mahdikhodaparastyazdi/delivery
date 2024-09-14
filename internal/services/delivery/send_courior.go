package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
	"delivery/internal/dto"
)

func (s Service) SendCourior(ctx context.Context, msg requests.SendCouriorRequest) error {
	var dtoMessage = dto.SendCourior{
		ProductID:           msg.ProductID,
		UserID:              msg.UserID,
		SourceLocation:      msg.SourceLocation,
		DestinationLocation: msg.DestinationLocation,
		StartTime:           msg.StartTime,
	}
	return s.queue3PL.Enqueue(dtoMessage)
}
