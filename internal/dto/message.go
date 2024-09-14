package dto

import "time"

type SendCourior struct {
	ProductID           int       `json:"product_id" validate:"required"`
	UserID              int       `json:"user_id" validate:"required"`
	SourceLocation      int       `json:"source_location" validate:"required"`
	DestinationLocation int       `json:"destination_location" validate:"required"`
	StartTime           time.Time `json:"start_time" validate:"required"`
}
type RecievedStatus struct {
	CouriorId int64  `json:"courior_id" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
