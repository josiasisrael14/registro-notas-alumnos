package section

import (
	"context"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type UseCase interface {
	Create(ctx context.Context, request model.Section) (model.ResponseStatusSection, error)
	GetWhere(ctx context.Context, id string) (model.Section, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Sections, error)
	Update(ctx context.Context, request model.Section) (model.ResponseStatusSection, error)
	//Delete(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}

type StorageSection interface {
	CreateSection(ctx context.Context, request model.Section) (model.ResponseStatusSection, error)
	GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Section, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Sections, error)
	UpdateSection(ctx context.Context, request model.Section) (model.ResponseStatusSection, error)
	//DeleteDegree(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}
