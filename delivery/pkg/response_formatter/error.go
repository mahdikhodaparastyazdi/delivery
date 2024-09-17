package response_formatter

import (
	"github.com/gin-gonic/gin"
)

func (r ResponseFormatter) Errors(c *gin.Context, m any, code int) {
	c.JSON(code, ResponseError{
		Errors: m,
	})
}

func (r ResponseFormatter) ErrorMessage(c *gin.Context, m string, code int) {
	c.JSON(code, ResponseError{
		ErrorMessage: m,
	})
}

func (r ResponseFormatter) ErrorsWithMessage(c *gin.Context, m any, message string, code int) {
	c.JSON(code, ResponseError{
		ErrorMessage: message,
		Errors:       m,
	})
}
