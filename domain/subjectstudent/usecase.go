package subjectstudent

import (
	"context"
	"fmt"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type SubjectStudent struct {
	storage StorageSubjectStudent
}

func New(storage StorageSubjectStudent) SubjectStudent {

	return SubjectStudent{
		storage: storage,
	}

}

func (s SubjectStudent) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.StudentSubjects, error) {
	ms, err := s.storage.GetAllWhere(ctx, specification)

	if err != nil {
		return nil, fmt.Errorf("subjectstudent.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}

func (s SubjectStudent) CreateSubjectStudent(ctx context.Context, request model.StudentSubject) (model.ResponseStatusSubjectStudent, error) {
	rs, err := s.storage.CreateSubjectStudent(ctx, request)

	if err != nil {
		return model.ResponseStatusSubjectStudent{}, fmt.Errorf("subjectstudent.storage.CreateSubjectStudent(): %w", err)
	}

	return rs, nil
}

func (s SubjectStudent) GetWhere(ctx context.Context, id string) (model.StudentSubject, error) {
	m, err := s.storage.GetWhere(ctx, repository.FieldsSpecification{Filters: repository.Fields{
		{Name: "student_subject_id", Value: id},
	}})
	if err != nil {
		return model.StudentSubject{}, fmt.Errorf("subjectstudent.storage.GetWhere(): %w", err)
	}

	return m, nil
}
