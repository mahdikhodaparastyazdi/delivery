package clientservice

import (
	"context"
	"delivery/internal/domain"
)

type clientRepository interface {
	GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error)
}

type Service struct {
	clientRepository clientRepository
}

func New(cr clientRepository) Service {
	return Service{
		clientRepository: cr,
	}
}
