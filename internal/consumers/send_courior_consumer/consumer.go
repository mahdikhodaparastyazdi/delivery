package courior_consumer

import (
	"bytes"
	"context"
	"delivery/internal/constants"
	"delivery/internal/domain"
	"delivery/internal/dto"
	"delivery/internal/services/delivery"
	"encoding/json"
	"time"
)

type send func(context.Context) error

func (c Consumer) Consume(ctx context.Context, message dto.SendCourior, retry, maxRetry int) error {
	if !delivery.IsInCurrentRange(time.Now(), message.StartTime) {
		return constants.ErrExpiryReached
	}
	if retry > maxRetry {
		nextTime := message.StartTime.Add(5 * time.Minute)
		err := c.queue3PL.Enqueue(message, &nextTime)
		if err != nil {
			return err
		}
		return constants.ErrBackOffRetry
	}
	couriorDomain := domain.COURIOR{
		ProductID:           message.ProductID,
		UserID:              message.UserID,
		SourceLocation:      message.SourceLocation,
		DestinationLocation: message.DestinationLocation,
		Status:              constants.COURIOR_STATUS_PENDING,
		StartTime:           message.StartTime,
	}
	courior, err := c.couriorRepo.Create(ctx, couriorDomain)
	if err != nil {
		return err
	}
	// TODO: selecting provider(3PL) can be conditional
	provider := c.resolver.ResolveCouriorProvider(constants.COURIOR_PROVIDER1)
	var bodyB = new(bytes.Buffer)
	err = json.NewEncoder(bodyB).Encode(courior)
	if err != nil {
		return constants.ErrUnexpected
	}
	return provider.SendCourior(ctx, "/courior", bodyB)
}
