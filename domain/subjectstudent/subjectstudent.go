package subjectstudent

import (
	"context"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type UseCase interface {
	//CreateSubjectTeacher(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error)
	//GetWhere(ctx context.Context, id string) (model.SubjectTeacher, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.StudentSubjects, error)
	//Update(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error)
	//Delete(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}
type StorageSubjectStudent interface {
	//GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.SubjectTeacher, error)
	//CreateSubjectTeacher(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.StudentSubjects, error)
	//UpdateSubjectTeacher(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error)
	//DeleteDegree(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}
