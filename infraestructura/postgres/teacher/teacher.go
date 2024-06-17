package teacher

import (
	"context"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/jackc/pgx/v5"
)

var _fieldWhere = []string{
	"id_teacher",
	"name",
	"surnames",
	"cellpone",
}

const _table = "teachers"

var (
	_psqlGetWhere    = repository.BuildSQLSelectFields(_table, _fieldWhere)
	_psqlGetAllWhere = repository.BuildSQLSelectFields(_table, _fieldWhere)
)

var (
	_psqlInsertTeacher = `INSERT INTO teachers(name,surnames,cellpone)VALUES($1,$2,$3)`
)

type Teacher struct {
	db model.PgxPool
}

func New(db model.PgxPool) Teacher {
	return Teacher{
		db: db,
	}
}

func (t Teacher) CreateTeacher(ctx context.Context, request model.Teacher) (model.ResponseStatusTeacher, error) {
	logTracer := register.NewPostgres(ctx, "postgres.teacher.CreateTeacher")

	logTracer.RegisterRequest(_psqlInsertTeacher, []any{request})

	_, err := t.db.Exec(ctx, _psqlInsertTeacher, request.Name, request.Surnames, request.CellPone)
	if err != nil {
		logTracer.RegisterFailed(err)
		return model.ResponseStatusTeacher{}, err
	}
	response := model.ResponseStatusTeacher{
		Response: "teacher create",
	}

	logTracer.RegisterResponse(response)

	return response, nil
}

func (t Teacher) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Teachers, error) {
	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAllWhere, specification)

	logTrace := register.NewPostgres(ctx, "postgres.teacher.GetAllWhere")
	logTrace.RegisterRequest(query, args)

	rows, err := t.db.Query(ctx, query, args...)
	if err != nil {
		logTrace.RegisterFailed(err)
		return nil, err
	}
	defer rows.Close()

	var ms model.Teachers
	for rows.Next() {
		m, err := t.scanRow(rows)
		if err != nil {
			logTrace.RegisterFailed(err)
			return nil, err
		}
		ms = append(ms, m)
	}

	logTrace.RegisterResponse(ms)

	return ms, nil
}

func (t Teacher) GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Teacher, error) {
	query, args := repository.BuildQueryAndArgs(_psqlGetWhere, specification)

	logTracer := register.NewPostgres(ctx, "postgres.teacher.GetWhere")
	logTracer.RegisterRequest(query, args)

	m, err := t.scanRow(t.db.QueryRow(ctx, query, args...))
	if err != nil {
		logTracer.RegisterFailed(err)
		return model.Teacher{}, err
	}

	logTracer.RegisterResponse(m)

	return m, nil
}

func (t Teacher) scanRow(p pgx.Row) (model.Teacher, error) {
	m := model.Teacher{}

	err := p.Scan(
		&m.IdTeacher,
		&m.Name,
		&m.Surnames,
		&m.CellPone,
	)
	if err != nil {
		return model.Teacher{}, err
	}

	return m, nil

}
