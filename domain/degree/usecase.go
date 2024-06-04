package degree

import (
	"context"
	"fmt"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type Degree struct {
	storage StorageDegree
}

func New(storage StorageDegree) Degree {

	return Degree{
		storage: storage,
	}

}

func (d Degree) Create(ctx context.Context, request model.Degree) (model.ResponseStatusDegree, error) {
	rs, err := d.storage.CreateDegree(ctx, request)

	if err != nil {
		return model.ResponseStatusDegree{}, fmt.Errorf("degree.storage.CreateDegree(): %w", err)
	}

	return rs, nil
}

func (d Degree) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Degrees, error) {
	ms, err := d.storage.GetAllWhere(ctx, specification)

	if err != nil {
		return nil, fmt.Errorf("degree.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}

func (d Degree) GetWhere(ctx context.Context, id string) (model.Degree, error) {
	m, err := d.storage.GetWhere(ctx, repository.FieldsSpecification{Filters: repository.Fields{
		{Name: "grades_id", Value: id},
	}})
	if err != nil {
		return model.Degree{}, fmt.Errorf("degree.storage.GetWhere(): %w", err)
	}

	return m, nil
}
