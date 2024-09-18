package courior_consumer

import (
	"bytes"
	"context"
	"delivery/internal/constants"
	"delivery/internal/domain"
	"delivery/internal/dto"
	"encoding/json"
	"errors"
	"time"
)

func (c Consumer) Consume(ctx context.Context, message dto.SendCourior, retry, maxRetry int) error {
	if time.Now().After(message.StartTime.Add(time.Hour * 2)) {
		return constants.ErrExpiryReached
	}
	if retry > maxRetry {
		nextTime := time.Now().Add(5 * time.Minute)
		err := c.queue3PL.Enqueue(message, &nextTime)
		if err != nil {
			return err
		}
		return constants.ErrBackOffRetry
	}
	couriorDomain := domain.Delivery{
		ProductID:           message.ProductID,
		UserID:              message.UserID,
		SourceLocation:      message.SourceLocation,
		DestinationLocation: message.DestinationLocation,
		Status:              constants.COURIOR_STATUS_PENDING,
		StartTime:           message.StartTime,
		CouriorID:           constants.COURIOR_PROVIDER1_ID,
		ProcessAt:           time.Now(),
	}
	courior, err := c.couriorRepo.Create(ctx, couriorDomain)

	// Provider needs to check and ignore repeated deliverId data
	if err != nil && !errors.Is(err, constants.ErrAlreadyExist) {
		return err
	}
	// TODO: selecting provider(3PL) can be conditional
	provider := c.resolver.ResolveCouriorProvider(constants.COURIOR_PROVIDER1_NAME)
	var bodyB = new(bytes.Buffer)
	err = json.NewEncoder(bodyB).Encode(courior)
	if err != nil {
		return constants.ErrUnexpected
	}
	return provider.SendCourior(ctx, "/courior", bodyB)
}
