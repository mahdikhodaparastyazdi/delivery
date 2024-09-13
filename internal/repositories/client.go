package repositories

import (
	"context"
	"delivery/internal/constants"
	"delivery/internal/domain"
	"delivery/internal/model"
	"errors"

	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return clientRepository{
		db: db,
	}
}

func (c clientRepository) GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error) {
	var client model.Client
	err := c.db.WithContext(ctx).Where("api_key = ? ", apiKey).Where("is_active = ?", true).Find(&client).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Client{}, constants.ErrWrongApiKey
		}
		return domain.Client{}, err
	}

	return client.ToDomain().(domain.Client), nil
}
func (c clientRepository) UpdateBalance(ctx context.Context, amount int, clientId uint) error {
	// TODO: Update Balance
	return nil
}
