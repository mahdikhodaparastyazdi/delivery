package repositories

import (
	"context"
	"delivery/internal/api/rest/requests"
	"delivery/internal/domain"
	"delivery/internal/model"

	"gorm.io/gorm"
)

type deliveryRepository struct {
	db *gorm.DB
}

func NewCouriorRepository(db *gorm.DB) DeliveryRepository {
	return deliveryRepository{
		db: db,
	}
}

func (s deliveryRepository) Create(ctx context.Context, courior domain.COURIOR) (domain.COURIOR, error) {
	mCourior := model.COURIOR{
		Mobile:     courior.Mobile,
		Content:    courior.Content,
		Status:     courior.Status.String(),
		ProviderId: courior.ProviderId,
		TemplateID: courior.TemplateID,
		ExpiredAt:  courior.ExpiresAt,
	}

	if err := s.db.WithContext(ctx).Create(&mCourior).Error; err != nil {
		return domain.COURIOR{}, err
	}
	ds := mCourior.ToDomain()

	return ds, nil
}
func (s deliveryRepository) SendCourior(ctx context.Context, msg requests.SendCouriorRequest) error {
	return nil
}
func (s deliveryRepository) ReceiveCouriorStatus(ctx context.Context, msg requests.CouriorStatusRequest) error {
	return nil
}
