package remissionguide

import (
	"context"
	"fmt"

	"notas/model"
)

type RemissionGuide struct {
	soap Soap
}

func New(soap Soap) RemissionGuide {
	return RemissionGuide{soap: soap}
}

func (r RemissionGuide) GetOrder(ctx context.Context, request model.RemissionOrderRequest) (model.RemissionOrderResponse, error) {
	ms, err := r.soap.Order(ctx, request)

	if err != nil {

		return model.RemissionOrderResponse{}, fmt.Errorf("order.soap.GetOrder():%w", err)
	}

	return ms, nil
}
