package courior_consumer

import (
	"context"
	"delivery/internal/constants"
	"delivery/internal/dto"
	"delivery/internal/services/delivery"
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
	provider, err := c.resolver.ResolveCouriorProvider(constants.COURIOR_PROVIDER1)
	if err != nil {
		return constants.ErrProviderNotFound
	}
	// TODO: insert row in courior table and send it
	return provider.SendCourior()
}
