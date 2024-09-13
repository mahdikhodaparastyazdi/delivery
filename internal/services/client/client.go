package clientservice

import (
	"context"
	"delivery/internal/domain"
)

func (s Service) GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error) {
	client, err := s.clientRepository.GetClientByApiKey(ctx, apiKey)
	if err != nil {
		return client, err
	}

	return client, nil
}
