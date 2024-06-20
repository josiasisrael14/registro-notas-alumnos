package section

import (
	"context"
	"fmt"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type Section struct {
	storage StorageSection
}

func New(storage StorageSection) Section {

	return Section{
		storage: storage,
	}

}

func (s Section) Create(ctx context.Context, request model.Section) (model.ResponseStatusSection, error) {
	rs, err := s.storage.CreateSection(ctx, request)

	if err != nil {
		return model.ResponseStatusSection{}, fmt.Errorf("section.storage.CreateDegree(): %w", err)
	}

	return rs, nil
}

func (s Section) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Sections, error) {
	ms, err := s.storage.GetAllWhere(ctx, specification)

	if err != nil {
		return nil, fmt.Errorf("section.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}

func (s Section) GetWhere(ctx context.Context, id string) (model.Section, error) {
	m, err := s.storage.GetWhere(ctx, repository.FieldsSpecification{Filters: repository.Fields{
		{Name: "section_id", Value: id},
	}})

	if err != nil {
		return model.Section{}, fmt.Errorf("section.storage.GetWhere(): %w", err)
	}

	return m, nil
}
func (s Section) Update(ctx context.Context, request model.Section) (model.ResponseStatusSection, error) {
	m, err := s.storage.UpdateSection(ctx, request)
	if err != nil {
		return model.ResponseStatusSection{}, fmt.Errorf("section.storage.UpdateSection():%w", err)
	}

	return m, nil
}
