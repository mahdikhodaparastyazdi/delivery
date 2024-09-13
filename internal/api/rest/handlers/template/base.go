package template

import (
	"context"
	"delivery/internal/api/rest/requests"
	"delivery/internal/api/rest/transformer"

	response_formatter "delivery/pkg/response_formatter"
)

type templateServiceInterface interface {
	SendTemplate(ctx context.Context, msg requests.SendWithTemplateRequest, clientId uint) error
}

type Handler struct {
	templateService     templateServiceInterface
	responseFormatter   response_formatter.ResponseFormatter
	templateTransformer transformer.TemplateTransformer
}

func New(templateService templateServiceInterface,
	responseFormatter response_formatter.ResponseFormatter,
	templateTrans transformer.TemplateTransformer,
) Handler {
	return Handler{
		templateService:     templateService,
		responseFormatter:   responseFormatter,
		templateTransformer: templateTrans,
	}
}
