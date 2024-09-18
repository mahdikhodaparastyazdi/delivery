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

func (s deliveryRepository) Create(ctx context.Context, courior domain.Delivery) (domain.Delivery, error) {
	mCourior := model.Delivery{
		CouriorID:           courior.CouriorID,
		ProductID:           courior.ProductID,
		UserID:              courior.UserID,
		SourceLocation:      courior.SourceLocation,
		DestinationLocation: courior.DestinationLocation,
		StartTime:           courior.StartTime,
		Status:              constants.COURIOR_STATUS_PENDING,
	}
	var existingCourior model.Delivery
	err := s.db.WithContext(ctx).Where("user_id = ? AND product_id = ? AND start_time = ?", courior.UserID,
		courior.ProductID, courior.StartTime).First(&existingCourior).Error
	if err == nil {
		return domain.Delivery{}, constants.ErrAlreadyExist
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Delivery{}, constants.ErrInternalServer
	}

	if err := s.db.WithContext(ctx).Create(&mCourior).Error; err != nil {
		return domain.Delivery{}, err
	}
	ds := mCourior.ToDomain()
	return ds, nil
}
func (s deliveryRepository) UpdateCouriorStatus(ctx context.Context,
	deliveryId uint,
	status constants.CouriorStatus) error {
	var deliver model.Delivery
	err := s.db.WithContext(ctx).Where("where id = ?", deliveryId).First(&deliver).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("delivery not found")
		}
		return fmt.Errorf("error finding delivery: %w", err)
	}
	deliver.Status = status
	if err := s.db.WithContext(ctx).Save(&deliver).Error; err != nil {
		return fmt.Errorf("error updating courior status: %w", err)
	}
	return nil
}
func (s deliveryRepository) GetById(ctx context.Context, deliveryID uint) (deliver domain.Delivery, err error) {
	err = s.db.WithContext(ctx).Where("where id = ?", deliveryID).First(&deliver).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Delivery{}, errors.New("delivery not found")
		}
		return domain.Delivery{}, fmt.Errorf("error finding delivery: %w", err)
	}
	return
}
