package template

import (
	"delivery/internal/api/rest/requests"
	"delivery/internal/constants"
	"errors"
	"net/http"

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
	err := h.deliveryService.SendCourior(c, req)
	if err == nil {
		h.responseFormatter.Success(c, nil, http.StatusOK)
		return
	}

	_ = c.Error(err)
	if errors.Is(err, constants.ErrTemplateNotFound) {
		h.responseFormatter.ErrorMessage(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if errors.Is(err, constants.ErrProviderIsNotActive) {
		h.responseFormatter.ErrorMessage(c, err.Error(), http.StatusBadRequest)
		return
	}
	if errors.Is(err, constants.ErrTemplateProviderNotDefined) {
		h.responseFormatter.ErrorMessage(c, err.Error(), http.StatusBadRequest)
		return
	}
	if errors.Is(err, constants.ErrExpiryDateTime) {
		h.responseFormatter.ErrorMessage(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	h.responseFormatter.ErrorMessage(c, constants.ErrInternalServer.Error(), constants.ErrInternalServerCode)
}
