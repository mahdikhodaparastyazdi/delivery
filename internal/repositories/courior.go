package repositories

import (
	"context"
	"delivery/internal/domain"
	"delivery/internal/model"

	"gorm.io/gorm"
)

type couriorRepository struct {
	db *gorm.DB
}

func NewCouriorRepository(db *gorm.DB) CouriorRepository {
	return couriorRepository{
		db: db,
	}
}

func (s couriorRepository) Create(ctx context.Context, courior domain.COURIOR) (domain.COURIOR, error) {
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
	ds := mCourior.ToDomain().(domain.COURIOR)

	return ds, nil
}
