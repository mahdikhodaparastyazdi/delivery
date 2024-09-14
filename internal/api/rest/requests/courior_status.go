package requests

type CouriorStatusRequest struct {
	CouriorId int64  `json:"courior_id" validate:"required"`
	Status    string `json:"status" validate:"required"`
}

func (s *CouriorStatusRequest) PrepareForValidation() {}
