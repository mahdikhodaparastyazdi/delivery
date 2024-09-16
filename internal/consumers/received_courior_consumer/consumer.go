package courior_consumer

import (
	"context"
	"delivery/internal/constants"
	"delivery/internal/dto"
	"time"
)

type send func(context.Context) error

func (c Consumer) Consume(ctx context.Context, message dto.RecievedStatus, retry, maxRetry int) error {
	if retry > maxRetry {
		nextTime := time.Now().Add(5 * time.Minute)
		err := c.queueCore.Enqueue(message, &nextTime)
		if err != nil {
			return err
		}
		return constants.ErrBackOffRetry
	}
	// TODO: request to core service
	return nil
}
