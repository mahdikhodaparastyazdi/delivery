package requests

import (
	"delivery/internal/constants"
)

type CouriorStatusRequest struct {
	CouriorId uint                    `json:"courior_id" validate:"required"`
	Status    constants.CouriorStatus `json:"status" validate:"required"`
}

func (s *CouriorStatusRequest) PrepareForValidation() {}
