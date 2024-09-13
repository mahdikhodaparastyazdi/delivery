package model

import (
	"delivery/internal/domain"
	"time"
)

type Client struct {
	ID          uint      `gorm:"column:id"`
	Balance     uint      `gorm:"column:balance"`
	COURIORRate string    `gorm:"column:courior_rate"`
	IsActive    bool      `gorm:"column:is_active"`
	Name        string    `gorm:"column:name"`
	ApiKey      string    `gorm:"api_key"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (c Client) TableName() string {
	return "clients"
}

func (c Client) ToDomain() domain.Domain {
	return domain.Client{
		ID:       c.ID,
		Username: c.Name,
		Token:    c.ApiKey,
		IsActive: c.IsActive,
	}
}
