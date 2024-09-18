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

// this handler only validate if time for sending courior not passed or greather than 4 days later and check
// start time hours between 9-23
func (h Handler) SendCourior(c *gin.Context) {
	var req requests.SendCouriorRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = c.Error(err)
		h.responseFormatter.ErrorMessage(c, constants.ErrValidation.Error(), constants.ErrValidationCode)
		return
	}
	//check hours between 9-23
	if !isValidTimeSlot(req.StartTime.Hour()) {
		h.responseFormatter.ErrorMessage(c, constants.ErrValidation.Error(), constants.ErrValidationCode)
		return
	}

	now := time.Now()
	lastValidHour := lastValidSlotHour(now)
	earlisetValidTime := calcEarliestValidTime(now, lastValidHour)
	fourDaysLater := now.Add(4 * 24 * time.Hour)

	if req.StartTime.Before(earlisetValidTime) || req.StartTime.After(fourDaysLater) {
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
func isValidTimeSlot(hour int) bool {
	if hour >= config.ValidTimeSlots[0] && hour < config.ValidTimeSlots[len(config.ValidTimeSlots)-1]+2 {
		return true
	}
	return false
}
func lastValidSlotHour(currentTime time.Time) int {
	currentHour := currentTime.Hour()
	for _, slot := range config.ValidTimeSlots {
		if slot <= currentHour && currentHour < slot+2 {
			return slot
		}
	}
	return config.ValidTimeSlots[len(config.ValidTimeSlots)-1]
}
func calcEarliestValidTime(now time.Time, lastValidHour int) (allowedTime time.Time) {
	timeToLastValidSlot := time.Duration(now.Hour()-lastValidHour) * time.Hour
	timeToLastValidSlot += time.Duration(now.Minute()) * time.Minute
	timeToLastValidSlot += time.Duration(now.Second()) * time.Second
	return now.Add(-timeToLastValidSlot)
}
