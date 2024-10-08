package model

import (
	"delivery/internal/constants"
	"delivery/internal/domain"
	"time"
)

type Delivery struct {
	ID                  uint                    `gorm:"column:id"`
	CouriorID           uint                    `gorm:"column:corior_id"`
	ProductID           uint                    `gorm:"column:product_id"`
	UserID              uint                    `gorm:"column:user_id"`
	SourceLocation      string                  `gorm:"column:source_location"`
	DestinationLocation string                  `gorm:"column:destination_location"`
	StartTime           time.Time               `gorm:"column:start_time"`
	Status              constants.CouriorStatus `gorm:"column:status"`
	CreatedAt           time.Time               `gorm:"column:created_at"`
	UpdatedAt           time.Time               `gorm:"column:updated_at"`
}

func (s Delivery) TableName() string {
	return "delivery"
}

func (s Delivery) ToDomain() domain.Delivery {
	return domain.Delivery{
		ID:                  s.ID,
		ProductID:           s.ProductID,
		UserID:              s.UserID,
		SourceLocation:      s.SourceLocation,
		DestinationLocation: s.DestinationLocation,
		StartTime:           s.StartTime,
		Status:              s.Status,
		CreatedAt:           s.CreatedAt,
		UpdatedAt:           s.UpdatedAt,
	}
}
