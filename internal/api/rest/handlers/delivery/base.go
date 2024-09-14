package delivery

import (
	"context"
	"delivery/internal/api/rest/requests"
	"delivery/internal/api/rest/transformer"

	response_formatter "delivery/pkg/response_formatter"
)

type deliveryServiceInterface interface {
	SendCourior(ctx context.Context, msg requests.SendCouriorRequest) error
	ReceiveCouriorStatus(ctx context.Context, msg requests.CouriorStatusRequest) error
}

type Handler struct {
	deliveryService     deliveryServiceInterface
	responseFormatter   response_formatter.ResponseFormatter
	templateTransformer transformer.DeliveryTransformer
}

func New(
	deliveryService deliveryServiceInterface,
	responseFormatter response_formatter.ResponseFormatter,
	deliveryTrans transformer.DeliveryTransformer,
) Handler {
	return Handler{
		deliveryService:     deliveryService,
		responseFormatter:   responseFormatter,
		templateTransformer: deliveryTrans,
	}
}
