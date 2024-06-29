package teacher

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"notas/model"
	"strings"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

const passwordLength = 12
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+[]{}|;:,.<>?/`~"

type Teacher struct {
	storage StorageTeacher
}

func New(storage StorageTeacher) Teacher {

	return Teacher{
		storage: storage,
	}

}

func (t Teacher) Create(ctx context.Context, request model.Teacher) (model.ResponseStatusTeacher, error) {
	nameUser := createUser(request.Name, request.Surnames)

	request.UserTeacher = nameUser

	password, err := createPassword()

	if err != nil {
		return model.ResponseStatusTeacher{
			Response: "error al generar password",
		}, err
	}

	request.Password = password

	rs, err := t.storage.CreateTeacher(ctx, request)

	if err != nil {
		return model.ResponseStatusTeacher{}, fmt.Errorf("section.storage.CreateDegree(): %w", err)
	}

	return rs, nil
}

func (t Teacher) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Teachers, error) {
	ms, err := t.storage.GetAllWhere(ctx, specification)

	if err != nil {
		return nil, fmt.Errorf("teacher.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}

func (t Teacher) GetWhere(ctx context.Context, id string) (model.Teacher, error) {
	m, err := t.storage.GetWhere(ctx, repository.FieldsSpecification{Filters: repository.Fields{
		{Name: "id_teacher", Value: id},
	}})
	if err != nil {
		return model.Teacher{}, fmt.Errorf("teacher.storage.GetWhere(): %w", err)
	}

	return m, nil
}

func (t Teacher) Update(ctx context.Context, request model.Teacher) (model.ResponseStatusTeacher, error) {
	m, err := t.storage.UpdateTeacher(ctx, request)
	if err != nil {
		return model.ResponseStatusTeacher{}, err
	}

	return m, nil
}

func createUser(name, surnames string) string {
	// Extraer el primer apellido
	parts := strings.Fields(surnames)
	if len(parts) > 0 {
		firstSurname := parts[0]                               // Tomar la primera parte como primer apellido
		nameUser := strings.ToLower(name + "_" + firstSurname) // Concatenar nombre y primer apellido
		return nameUser
	}

	// Si no hay apellidos, usar solo el nombre
	return strings.ToLower(name)

}

func createPassword() (string, error) {
	password := make([]byte, passwordLength)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}
	return string(password), nil
}
