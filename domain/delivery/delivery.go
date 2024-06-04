package delivery

import (
	"context"

	"notas/model"
)

type UseCase interface {
	delivery(ctx context.Context, request model.DeliveryRequest) (model.DeliveryResponse, error)
}

type Soap interface {
	delivery(ctx context.Context, request model.DeliveryRequest) (model.DeliveryResponse, error)
}
