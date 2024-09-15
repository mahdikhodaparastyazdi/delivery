package courior_consumer

import (
	"context"
	"delivery/internal/dto"
)

type send func(context.Context) error

func (c Consumer) Consume(ctx context.Context, message dto.RecievedStatus, retry, maxRetry int) error {
	return nil
}
