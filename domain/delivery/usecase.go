package delivery

import (
	"context"
	"fmt"

	"notas/model"
)

type Delivery struct {
	soap Soap
}

func New(soap Soap) Delivery {
	return Delivery{soap: soap}
}

func (d Delivery) delivery(ctx context.Context, request model.DeliveryRequest) (model.DeliveryResponse, error) {
	ms, err := d.soap.delivery(ctx, request)

	if err != nil {
		return model.DeliveryResponse{}, fmt.Errorf("delivery.soap.delivery():%w", err)
	}

	return ms, nil
}
