package teacher

import (
	"context"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type UseCase interface {
	Create(ctx context.Context, request model.Teacher) (model.ResponseStatusTeacher, error)
	GetWhere(ctx context.Context, id string) (model.Teacher, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Teachers, error)
	Update(ctx context.Context, request model.Teacher) (model.ResponseStatusTeacher, error)
	//Delete(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}

type StorageTeacher interface {
	CreateTeacher(ctx context.Context, request model.Teacher) (model.ResponseStatusTeacher, error)
	GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Teacher, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Teachers, error)
	UpdateTeacher(ctx context.Context, request model.Teacher) (model.ResponseStatusTeacher, error)
	//DeleteDegree(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}
