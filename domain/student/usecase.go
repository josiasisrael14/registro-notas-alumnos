package student

import (
	"context"
	"fmt"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"

	"notas/model"
)

type Student struct {
	storage StorageStudent
}

func New(storage StorageStudent) Student {

	return Student{
		storage: storage,
	}

}

func (s Student) Create(ctx context.Context, request model.Student) (model.ResponseStatusStudent, error) {
	m, err := s.storage.CreateStudent(ctx, request)
	if err != nil {
		return model.ResponseStatusStudent{}, fmt.Errorf("student.storage.CreateStudent():%w", err)
	}

	return m, nil
}

func (s Student) GetWhere(ctx context.Context, id string) (model.Student, error) {
	m, err := s.storage.GetWhere(ctx, repository.FieldsSpecification{Filters: repository.Fields{
		{Name: "student_id", Value: id},
	}})
	if err != nil {
		return model.Student{}, fmt.Errorf("student.storage.GetWhere(): %w", err)
	}

	return m, nil
}

func (s Student) Update(ctx context.Context, request model.Student) (model.ResponseStatusStudent, error) {
	m, err := s.storage.UpdateStudent(ctx, request)
	if err != nil {
		return model.ResponseStatusStudent{}, fmt.Errorf("student.storage.UpdateStudent():%w", err)
	}

	return m, nil
}

func (s Student) Delete(ctx context.Context, id string) (model.ResponseStatusStudent, error) {
	m, err := s.storage.DeleteStudent(ctx, id)
	if err != nil {
		return model.ResponseStatusStudent{}, fmt.Errorf("student.storage.DeleteStudent():%w", err)
	}

	return m, nil
}

func (s Student) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Students, error) {
	ms, err := s.storage.GetAllWhere(ctx, specification)

	if err != nil {
		return nil, fmt.Errorf("student.storage.GetAllWhere(): %w", err)
	}

	/*for i := range ms {
		year, month, day := ms[i].BirthDate.Date()                          // Extraer año, mes y día
		ms[i].BirthDate = time.Date(year, month, day, 0, 0, 0, 0, time.UTC) // Establecer hora a cero

		// Formatear la fecha sin ceros iniciales
		ms[i].BirthDate.Format("2006-01-02")                                           // Formatear a YYYY-MM-DD
		strings.Trim(strings.TrimLeft(ms[i].BirthDate.Format("2006-01-02"), "0"), "-") // Eliminar ceros iniciales y guiones

	}*/

	// Convertir las fechas a Date

	return ms, nil
}
