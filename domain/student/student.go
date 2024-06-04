package student

import (
	"context"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"

	"notas/model"
)

type UseCase interface {
	Create(ctx context.Context, request model.Student) (model.ResponseStatusStudent, error)
	GetWhere(ctx context.Context, id string) (model.Student, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Students, error)
	Update(ctx context.Context, request model.Student) (model.ResponseStatusStudent, error)
	Delete(ctx context.Context, id string) (model.ResponseStatusStudent, error)
}

type StorageStudent interface {
	CreateStudent(ctx context.Context, request model.Student) (model.ResponseStatusStudent, error)
	GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Student, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Students, error)
	UpdateStudent(ctx context.Context, request model.Student) (model.ResponseStatusStudent, error)
	DeleteStudent(ctx context.Context, id string) (model.ResponseStatusStudent, error)
}
