package stuff

import (
	"context"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"

	"notas/model"
)

type UseCase interface {
	Create(ctx context.Context, request model.Stuff) (model.ResponseStatusStuff, error)
	GetWhere(ctx context.Context, id string) (model.Stuff, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Stuffs, error)
	Update(ctx context.Context, request model.Stuff) (model.ResponseStatusStuff, error)
	Delete(ctx context.Context, id string) (model.ResponseStatusStuff, error)
}

type StorageStuff interface {
	CreateStuff(ctx context.Context, request model.Stuff) (model.ResponseStatusStuff, error)
	GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Stuff, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Stuffs, error)
	UpdateStuff(ctx context.Context, request model.Stuff) (model.ResponseStatusStuff, error)
	DeleteStuff(ctx context.Context, id string) (model.ResponseStatusStuff, error)
}
