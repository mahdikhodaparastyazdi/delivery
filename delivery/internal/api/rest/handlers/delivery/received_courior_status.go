package delivery

import (
	"delivery/internal/api/rest/requests"
	"delivery/internal/constants"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h Handler) ReceiveCouriorStatus(c *gin.Context) {
	var req requests.CouriorStatusRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = c.Error(err)
		h.responseFormatter.ErrorMessage(c, constants.ErrValidation.Error(), constants.ErrValidationCode)
		return
	}
	err := h.deliveryService.ReceiveCouriorStatus(c, req)
	if err == nil {
		h.responseFormatter.Success(c, nil, http.StatusOK)
		return
	}
	_ = c.Error(err)
	h.responseFormatter.ErrorMessage(c, constants.ErrInternalServer.Error(), constants.ErrInternalServerCode)
}
