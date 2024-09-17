package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
	"delivery/internal/dto"
	"time"
)

func (s Service) SendCourior(ctx context.Context, msg requests.SendCouriorRequest, couriorSendTime time.Time) error {
	processTime := couriorSendTime.Add(-time.Hour * 1)
	var dtoMessage = dto.SendCourior{
		ProductID:           msg.ProductID,
		UserID:              msg.UserID,
		SourceLocation:      msg.SourceLocation,
		DestinationLocation: msg.DestinationLocation,
		StartTime:           msg.StartTime,
		ProcessAt:           processTime,
	}
	if msg.StartTime.Before(couriorSendTime.Add(time.Hour * 2)) {
		return s.queue3PL.Enqueue(dtoMessage, nil)
	}
	return s.queue3PL.Enqueue(dtoMessage, &processTime)
}
