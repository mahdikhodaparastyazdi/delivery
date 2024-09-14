package requests

import "time"

type SendCouriorRequest struct {
	ProductID           int       `json:"product_id" validate:"required"`
	UserID              int       `json:"user_id" validate:"required"`
	SourceLocation      int       `json:"source_location" validate:"required"`
	DestinationLocation int       `json:"destination_location" validate:"required"`
	StartTime           time.Time `json:"start_time" validate:"required"`
}

func (s *SendCouriorRequest) PrepareForValidation() {}
