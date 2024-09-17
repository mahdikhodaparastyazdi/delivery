package provider1

import (
	"bytes"
	"context"
	"delivery/internal/constants"
	"net/http"
)

func (p provider1) SendCourior(ctx context.Context, path string, buf *bytes.Buffer) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseUrl3PL+path, buf)
	if err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(apiKeyHeader, p.apiKey3Pl)

	response, err := p.hc.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return constants.ErrWrongStatus
	}
	return nil
}
