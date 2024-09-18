package requests

import (
	"delivery/internal/constants"
)

type CouriorStatusRequest struct {
	DeliverId uint                    `json:"deliver_id" validate:"required"`
	Status    constants.CouriorStatus `json:"status" validate:"required"`
}

func (s *CouriorStatusRequest) PrepareForValidation() {}
