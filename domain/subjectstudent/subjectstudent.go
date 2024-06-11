package subjectstudent

import (
	"context"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
)

type UseCase interface {
	CreateSubjectStudent(ctx context.Context, request model.StudentSubject) (model.ResponseStatusSubjectStudent, error)
	GetWhere(ctx context.Context, id string) (model.StudentSubject, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.StudentSubjects, error)
	//Update(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error)
	//Delete(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}

type StorageSubjectStudent interface {
	GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.StudentSubject, error)
	CreateSubjectStudent(ctx context.Context, request model.StudentSubject) (model.ResponseStatusSubjectStudent, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.StudentSubjects, error)
	//UpdateSubjectTeacher(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error)
	//DeleteDegree(ctx context.Context, id string) (model.ResponseStatusDegree, error)
}

/*type StorageStudent interface {
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Students, error)
}*/
