package domain

import (
	"delivery/internal/constants"
	"time"
)

type COURIOR struct {
	ID                  uint
	ProductID           uint
	UserID              uint
	SourceLocation      string
	DestinationLocation string
	StartTime           time.Time
	Status              constants.CouriorStatus
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
