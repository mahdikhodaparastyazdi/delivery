package domain

import (
	"delivery/internal/constants"
	"time"
)

type Delivery struct {
	ID                  uint
	ProductID           uint
	CouriorID           uint
	UserID              uint
	SourceLocation      string
	DestinationLocation string
	StartTime           time.Time
	Status              constants.CouriorStatus
	CreatedAt           time.Time
	UpdatedAt           time.Time
	ProcessAt           time.Time
}
