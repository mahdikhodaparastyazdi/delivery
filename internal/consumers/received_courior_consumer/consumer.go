package courior_consumer

import (
	"bytes"
	"context"
	"delivery/internal/constants"
	"delivery/internal/dto"
	"encoding/json"
	"net/http"
	"time"
)

type send func(context.Context) error

func (c Consumer) Consume(ctx context.Context, message dto.RecievedStatus, retry, maxRetry int) error {
	if time.Now().Hour() >= 23 || time.Now().Hour() <= 8 {
		return constants.ErrExpiryReached
	}
	if retry > maxRetry {
		nextTime := time.Now().Add(5 * time.Minute)
		err := c.queueCore.Enqueue(message, &nextTime)
		if err != nil {
			return err
		}
		return constants.ErrBackOffRetry
	}
	err := c.couriorRepository.UpdateCouriorStatus(ctx, message.CouriorId, message.Status)
	if err != nil {
		return err
	}
	var bodyB = new(bytes.Buffer)
	err = json.NewEncoder(bodyB).Encode(message)
	return c.requestCore(ctx, "/webhook", bodyB)
}
func (c Consumer) requestCore(ctx context.Context, path string, buf *bytes.Buffer) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseUrlCore+path, buf)
	if err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(apiKeyHeader, c.apiKeyCore)

	response, err := c.hc.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return constants.ErrWrongStatus
	}
	return nil
}
