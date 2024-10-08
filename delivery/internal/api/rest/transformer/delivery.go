package transformer

import (
	"context"
	"delivery/internal/domain"
)

type DeliveryTransformer struct{}

type ResponseDelivery struct {
	ID                  uint   `json:"id" example:"1"`
	Code                string `json:"code" example:"42"`
	Content             string `json:"content"`
	Params              string `json:"params" example:"username,age"`
	UseProviderTemplate bool   `json:"use_provider_template" example:"true"`
	ActiveProviderID    uint   `json:"active_provider_id" example:"4"`
	Type                string `json:"type" example:"OTP"`
}

func NewDeliveryTransformer() DeliveryTransformer {
	return DeliveryTransformer{}
}

func (t DeliveryTransformer) Transform(_ context.Context, tmpl domain.Delivery) ResponseDelivery {
	var resp ResponseDelivery

	return resp
}

func (t DeliveryTransformer) TransformMany(ctx context.Context, templates []domain.Delivery) []ResponseDelivery {
	var result []ResponseDelivery
	for _, tm := range templates {
		transProvider := t.Transform(ctx, tm)

		result = append(result, transProvider)
	}

	return result
}
