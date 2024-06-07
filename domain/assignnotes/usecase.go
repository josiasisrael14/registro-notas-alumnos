package assignnotes

import (
	"context"
	"fmt"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type AssignNotes struct {
	storage StorageAssignNote
}

func New(storage StorageAssignNote) AssignNotes {

	return AssignNotes{
		storage: storage,
	}

}

func (a AssignNotes) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.AssignNotes, error) {
	ms, err := a.storage.GetAllWhere(ctx, specification)

	if err != nil {
		return nil, fmt.Errorf("assignnotes.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}
