package teacher

import (
	"context"
	"fmt"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type Teacher struct {
	storage StorageTeacher
}

func New(storage StorageTeacher) Teacher {

	return Teacher{
		storage: storage,
	}

}

func (t Teacher) Create(ctx context.Context, request model.Teacher) (model.ResponseStatusTeacher, error) {
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
