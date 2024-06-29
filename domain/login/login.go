package login

import (
	"context"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type UseCase interface {
	Login(ctx context.Context, request model.Login) (model.ResponseStatusLogin, error)
}

type StorageLogin interface {
	LoginAcceso(ctx context.Context, specification repository.FieldsSpecification) (model.ResponseStatusLogin, error)
}
