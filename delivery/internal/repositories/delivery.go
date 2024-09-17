package repositories

import (
	"context"
	"delivery/internal/constants"
	"delivery/internal/domain"
	"delivery/internal/model"
	"errors"
	"fmt"

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
		CouriorID:           courior.CouriorID,
		ProductID:           courior.ProductID,
		UserID:              courior.UserID,
		SourceLocation:      courior.SourceLocation,
		DestinationLocation: courior.DestinationLocation,
		StartTime:           courior.StartTime,
		Status:              constants.COURIOR_STATUS_PENDING,
	}
	var existingCourior model.COURIOR
	err := s.db.WithContext(ctx).Where("user_id = ? AND product_id = ? AND start_time = ?", courior.UserID,
		courior.ProductID, courior.StartTime).First(&existingCourior).Error
	if err == nil {
		return domain.COURIOR{}, constants.ErrAlreadyExist
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.COURIOR{}, constants.ErrInternalServer
	}

	if err := s.db.WithContext(ctx).Create(&mCourior).Error; err != nil {
		return domain.COURIOR{}, err
	}
	ds := mCourior.ToDomain()
	return ds, nil
}
func (s deliveryRepository) UpdateCouriorStatus(ctx context.Context,
	couriorId uint,
	status constants.CouriorStatus) error {
	var courior model.COURIOR
	result := s.db.WithContext(ctx).First(&courior, couriorId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("courior not found")
		}
		return fmt.Errorf("error finding courior: %w", result.Error)
	}
	courior.Status = status
	if err := s.db.WithContext(ctx).Save(&courior).Error; err != nil {
		return fmt.Errorf("error updating courior status: %w", err)
	}
	return nil
}
