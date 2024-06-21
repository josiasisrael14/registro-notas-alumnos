package student

import (
	"context"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/jackc/pgx/v5"

	"notas/model"
)

var (
	_psqlInsertStudent = `INSERT INTO students( name,surnames,birth_date )VALUES($1,$2,$3)`
)

var (
	_psqlUpdateStudent = `update students set name=$1,surnames=$2,birth_date=$3 where student_id=$4`
)

var (
	_psqlDeleteStuff = `delete from students where student_id=$1`
)

var _fieldWhere = []string{
	"student_id",
	"name",
	"surnames",
	"birth_date",
}

const _table = "students"

var (
	_psqlGetWhere    = repository.BuildSQLSelectFields(_table, _fieldWhere)
	_psqlGetAllWhere = repository.BuildSQLSelectFields(_table, _fieldWhere)
)

type Student struct {
	db model.PgxPool
}

func New(db model.PgxPool) Student {
	return Student{
		db: db,
	}
}

func (s Student) CreateStudent(ctx context.Context, request model.Student) (model.ResponseStatusStudent, error) {
	logTracer := register.NewPostgres(ctx, "postgres.student.CreateStudent")

	logTracer.RegisterRequest(_psqlInsertStudent, []any{request})

	_, err := s.db.Exec(ctx, _psqlInsertStudent, request.NameStudent, request.LastName, request.BirthDate.Format(model.DateFormat)) //Format(model.DateFormat))
	if err != nil {
		logTracer.RegisterFailed(err)
		return model.ResponseStatusStudent{}, err
	}
	response := model.ResponseStatusStudent{

		Response: "students create success",
	}

	logTracer.RegisterResponse(response)

	return response, nil
}

func (s Student) GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Student, error) {
	query, args := repository.BuildQueryAndArgs(_psqlGetWhere, specification)

	logTracer := register.NewPostgres(ctx, "postgres.student.GetWhere")
	logTracer.RegisterRequest(query, args)

	m, err := s.scanRow(s.db.QueryRow(ctx, query, args...))
	if err != nil {
		logTracer.RegisterFailed(err)
		return model.Student{}, err
	}

	logTracer.RegisterResponse(m)

	return m, nil
}

func (s Student) UpdateStudent(ctx context.Context, request model.Student) (model.ResponseStatusStudent, error) {
	logTrace := register.NewPostgres(ctx, "postgres.student.UpdateStudent")

	logTrace.RegisterRequest(_psqlUpdateStudent, []any{request})

	_, err := s.db.Exec(ctx, _psqlUpdateStudent, request.NameStudent, request.LastName, request.BirthDate.Format(model.DateFormat), request.StudentId)

	if err != nil {
		logTrace.RegisterFailed(err)
		return model.ResponseStatusStudent{}, err
	}

	response := model.ResponseStatusStudent{

		Response: "update success",
	}

	logTrace.RegisterResponse(response)

	return response, nil
}

func (s Student) DeleteStudent(ctx context.Context, id string) (model.ResponseStatusStudent, error) {
	logTrace := register.NewPostgres(ctx, "postgres.student.DeleteStudent")

	logTrace.RegisterRequest(_psqlDeleteStuff, []any{id})

	_, err := s.db.Exec(ctx, _psqlDeleteStuff, id)

	if err != nil {
		logTrace.RegisterFailed(err)
		return model.ResponseStatusStudent{}, err
	}

	response := model.ResponseStatusStudent{

		Response: "delete success",
	}

	logTrace.RegisterResponse(response)

	return response, nil
}

func (s Student) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Students, error) {
	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAllWhere, specification)

	logTrace := register.NewPostgres(ctx, "postgres.student.GetAllWhere")
	logTrace.RegisterRequest(query, args)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		logTrace.RegisterFailed(err)
		return nil, err
	}
	defer rows.Close()

	var ms model.Students
	for rows.Next() {
		m, err := s.scanRow(rows)
		if err != nil {
			logTrace.RegisterFailed(err)
			return nil, err
		}
		ms = append(ms, m)
	}

	logTrace.RegisterResponse(ms)

	return ms, nil
}

func (s Student) scanRow(p pgx.Row) (model.Student, error) {
	m := model.Student{}
	err := p.Scan(
		&m.StudentId,
		&m.NameStudent,
		&m.LastName,
		&m.BirthDate,
	)
	if err != nil {
		return model.Student{}, err
	}
	return m, nil

}
