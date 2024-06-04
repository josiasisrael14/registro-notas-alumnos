package stuff

import (
	"context"
	"fmt"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"

	"notas/model"
)

type Stuff struct {
	storage StorageStuff
}

func New(storage StorageStuff) Stuff {

	return Stuff{
		storage: storage,
	}

}

func (s Stuff) Create(ctx context.Context, request model.Stuff) (model.ResponseStatusStuff, error) {

	rs, err := s.storage.CreateStuff(ctx, request)

	if err != nil {
		return model.ResponseStatusStuff{}, fmt.Errorf("stuff.storage.CreateStuff(): %w", err)
	}

	return rs, nil

}

func (s Stuff) GetWhere(ctx context.Context, id string) (model.Stuff, error) {
	m, err := s.storage.GetWhere(ctx, repository.FieldsSpecification{Filters: repository.Fields{
		{Name: "materiaid", Value: id},
	}})
	if err != nil {
		return model.Stuff{}, fmt.Errorf("stuff.storage.GetWhere(): %w", err)
	}

	return m, nil
}

func (s Stuff) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Stuffs, error) {
	m, err := s.storage.GetAllWhere(ctx, specification)
	if err != nil {
		return nil, fmt.Errorf("stuff.storage.GetAllWhere():%w", err)
	}

	return m, nil
}

func (s Stuff) Update(ctx context.Context, request model.Stuff) (model.ResponseStatusStuff, error) {
	m, err := s.storage.UpdateStuff(ctx, request)
	if err != nil {
		return model.ResponseStatusStuff{}, err
	}

	return m, nil
}

func (s Stuff) Delete(ctx context.Context, id string) (model.ResponseStatusStuff, error) {
	m, err := s.storage.DeleteStuff(ctx, id)
	if err != nil {
		return model.ResponseStatusStuff{}, err
	}

	return m, nil
}
