package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
	"delivery/internal/config"
	"delivery/internal/dto"
	"time"
)

func (s Service) SendCourior(ctx context.Context, msg requests.SendCouriorRequest, now time.Time) error {
	var dtoMessage = dto.SendCourior{
		ProductID:           msg.ProductID,
		UserID:              msg.UserID,
		SourceLocation:      msg.SourceLocation,
		DestinationLocation: msg.DestinationLocation,
		StartTime:           msg.StartTime,
	}
	if IsInCurrentRange(now, msg.StartTime) {
		err := s.queue3PL.Enqueue(dtoMessage, nil)
		if err != nil {
			return err
		}
	}
	return s.queue3PL.Enqueue(dtoMessage, &msg.StartTime)
}

func IsInCurrentRange(now, timestamp time.Time) bool {
	nowHour := now.Hour()
	timestampHour := timestamp.Hour()

	for _, validHour := range config.ValidTimeSlots {
		if nowHour >= validHour && nowHour < validHour+2 && timestampHour == validHour {
			return true
		}
	}
	return false
}
