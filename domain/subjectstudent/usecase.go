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
