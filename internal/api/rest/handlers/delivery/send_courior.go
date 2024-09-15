package delivery

import (
	"delivery/internal/api/rest/requests"
	"delivery/internal/config"
	"delivery/internal/constants"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h Handler) SendCourior(c *gin.Context) {
	var req requests.SendCouriorRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = c.Error(err)
		h.responseFormatter.ErrorMessage(c, constants.ErrValidation.Error(), constants.ErrValidationCode)
		return
	}
	now := time.Now()
	fourDaysLater := now.Add(4 * 24 * time.Hour)

	// Ensure the provided startTime is between now and 4 days later
	if req.StartTime.Before(now) || req.StartTime.After(fourDaysLater) {
		h.responseFormatter.ErrorMessage(c, "StartTime must be between now and 4 days later", http.StatusBadRequest)
		return
	}
	err := h.deliveryService.SendCourior(c, req, now)
	if err == nil {
		h.responseFormatter.Success(c, nil, http.StatusOK)
		return
	}
	_ = c.Error(err)
	h.responseFormatter.ErrorMessage(c, constants.ErrInternalServer.Error(), constants.ErrInternalServerCode)
}

// TODO: need bring to validator package
func ValidateTime(t time.Time) bool {
	hour := t.Hour()
	for _, validHour := range config.ValidTimeSlots {
		if hour == validHour && t.Minute() == 0 && t.Second() == 0 {
			return true
		}
	}
	return false
}
