package remissionguide

import (
	"context"

	"notas/model"
)

type UseCase interface {
	GetOrder(ctx context.Context, request model.RemissionOrderRequest) (model.RemissionOrderResponse, error)
}

type Soap interface {
	Order(ctx context.Context, request model.RemissionOrderRequest) (model.RemissionOrderResponse, error)
}
