package requests

import "time"

type SendCouriorRequest struct {
	ProductID           uint      `json:"product_id" validate:"required"`
	UserID              uint      `json:"user_id" validate:"required"`
	SourceLocation      string    `json:"source_location" validate:"required"`
	DestinationLocation string    `json:"destination_location" validate:"required"`
	StartTime           time.Time `json:"start_time" validate:"required"`
}

func (s *SendCouriorRequest) PrepareForValidation() {}
