package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
	"delivery/internal/dto"
	"time"
)

func (s Service) SendCourior(ctx context.Context,
	msg requests.SendCouriorRequest,
	now time.Time) error {
	var dtoMessage = dto.SendCourior{
		ProductID:           msg.ProductID,
		UserID:              msg.UserID,
		SourceLocation:      msg.SourceLocation,
		DestinationLocation: msg.DestinationLocation,
		StartTime:           msg.StartTime,
	}
	if msg.StartTime.Before(now.Add(time.Hour * 1)) {
		return s.queue3PL.Enqueue(dtoMessage, nil)
	}
	processTime := dtoMessage.StartTime.Add(time.Hour * 1)
	return s.queue3PL.Enqueue(dtoMessage, &processTime)
}
