package login

import (
	"context"
	"fmt"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type Login struct {
	storage StorageLogin
}

func New(storage StorageLogin) Login {

	return Login{
		storage: storage,
	}

}

func (l Login) Login(ctx context.Context, request model.Login) (model.ResponseStatusLogin, error) {
	m, err := l.storage.LoginAcceso(ctx, repository.FieldsSpecification{Filters: repository.Fields{
		{Name: "login", Value: request.NameUser},
		{Name: "password", Value: request.Password},
	}})

	if err != nil {
		return model.ResponseStatusLogin{}, fmt.Errorf("login.storage.LoginAcceso(): %w", err)
	}

	return m, nil
}
