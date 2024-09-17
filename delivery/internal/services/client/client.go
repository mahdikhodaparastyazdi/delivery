package clientservice

import (
	"context"
	"delivery/internal/domain"
)

func (s Service) GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error) {
	return s.clientRepository.GetClientByApiKey(ctx, apiKey)
}
