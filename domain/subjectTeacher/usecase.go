package subjectTeacher

import (
	"context"
	"fmt"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type SubjectTeacher struct {
	storage StorageSubjectTeacher
}

func New(storage StorageSubjectTeacher) SubjectTeacher {

	return SubjectTeacher{
		storage: storage,
	}

}

func (s SubjectTeacher) CreateSubjectTeacher(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error) {

	rs, err := s.storage.CreateSubjectTeacher(ctx, request)

	if err != nil {
		return model.ResponseStatusSubjectTeacher{}, fmt.Errorf("subjectTeacher.storage.CreateSubjectTeacher(): %w", err)
	}

	return rs, nil
}

func (s SubjectTeacher) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.SubjectTeachers, error) {
	ms, err := s.storage.GetAllWhere(ctx, specification)

	if err != nil {
		return nil, fmt.Errorf("subjectTeacher.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}
func (s SubjectTeacher) Update(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error) {
	rs, err := s.storage.UpdateSubjectTeacher(ctx, request)

	if err != nil {
		return model.ResponseStatusSubjectTeacher{}, fmt.Errorf("subjectTeacher.storage.UpdateSubjectTeacher(): %w", err)
	}

	return rs, nil

}

func (s SubjectTeacher) GetWhere(ctx context.Context, id string) (model.SubjectTeacher, error) {
	m, err := s.storage.GetWhere(ctx, repository.FieldsSpecification{Filters: repository.Fields{
		{Name: "subjects_teacher_id", Value: id},
	}})
	if err != nil {
		return model.SubjectTeacher{}, fmt.Errorf("subjectTeacher.storage.GetWhere(): %w", err)
	}

	return m, nil
}
