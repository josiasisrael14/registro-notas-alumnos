package degree

import (
	"context"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type UseCase interface {
	Create(ctx context.Context, request model.Degree) (model.ResponseStatusDegree, error)
	GetWhere(ctx context.Context, id string) (model.Degree, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Degrees, error)
	//Update(ctx context.Context, request model.Degree) (model.ResponseStatusDegree, error)
	//Delete(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}

type StorageDegree interface {
	CreateDegree(ctx context.Context, request model.Degree) (model.ResponseStatusDegree, error)
	GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Degree, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Degrees, error)
	//UpdateDegree(ctx context.Context, request model.Degree) (model.ResponseStatusDegree, error)
	//DeleteDegree(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}
