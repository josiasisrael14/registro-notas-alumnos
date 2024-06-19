package degree

import (
	"context"
	"database/sql"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/jackc/pgx/v5"

	"notas/model"
)

var _fieldWhere = []string{
	"grades_id",
	"grade_name",
	"specific_level",
}

const _table = "grades"

var (
	_psqlGetWhere    = repository.BuildSQLSelectFields(_table, _fieldWhere)
	_psqlGetAllWhere = repository.BuildSQLSelectFields(_table, _fieldWhere)
)

var (
	_psqlInsertDegree = `INSERT INTO grades( grade_name,specific_level )VALUES($1,$2)`
)

var (
	_psqlUpdateDegree = `update grades set grade_name=$1,specific_level=$2 where grades_id=$3`
)

type Degree struct {
	db model.PgxPool
}

func New(db model.PgxPool) Degree {
	return Degree{
		db: db,
	}
}

func (d Degree) CreateDegree(ctx context.Context, request model.Degree) (model.ResponseStatusDegree, error) {
	logTracer := register.NewPostgres(ctx, "postgres.degree.CreateDegree")

	logTracer.RegisterRequest(_psqlInsertDegree, []any{request})

	_, err := d.db.Exec(ctx, _psqlInsertDegree, request.NameDegree, request.SpecificLevel)
	if err != nil {
		logTracer.RegisterFailed(err)
		return model.ResponseStatusDegree{}, err
	}
	response := model.ResponseStatusDegree{
		Response: "Degree create",
	}

	logTracer.RegisterResponse(response)

	return response, nil
}

func (d Degree) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Degrees, error) {
	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAllWhere, specification)

	logTrace := register.NewPostgres(ctx, "postgres.degree.GetAllWhere")
	logTrace.RegisterRequest(query, args)

	rows, err := d.db.Query(ctx, query, args...)
	if err != nil {
		logTrace.RegisterFailed(err)
		return nil, err
	}
	defer rows.Close()

	var ms model.Degrees
	for rows.Next() {
		m, err := d.scanRow(rows)
		if err != nil {
			logTrace.RegisterFailed(err)
			return nil, err
		}
		ms = append(ms, m)
	}

	logTrace.RegisterResponse(ms)

	return ms, nil
}

func (d Degree) GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Degree, error) {
	query, args := repository.BuildQueryAndArgs(_psqlGetWhere, specification)

	logTracer := register.NewPostgres(ctx, "postgres.degree.GetWhere")
	logTracer.RegisterRequest(query, args)

	m, err := d.scanRow(d.db.QueryRow(ctx, query, args...))
	if err != nil {
		logTracer.RegisterFailed(err)
		return model.Degree{}, err
	}

	logTracer.RegisterResponse(m)

	return m, nil
}

func (d Degree) UpdateDegree(ctx context.Context, request model.Degree) (model.ResponseStatusDegree, error) {
	logTrace := register.NewPostgres(ctx, "postgres.degree.UpdateStuff")

	logTrace.RegisterRequest(_psqlUpdateDegree, []any{request})

	_, err := d.db.Exec(ctx, _psqlUpdateDegree, request.NameDegree, request.SpecificLevel, request.IdDegree)

	if err != nil {
		logTrace.RegisterFailed(err)
		return model.ResponseStatusDegree{}, err
	}

	response := model.ResponseStatusDegree{

		Response: "degree success",
	}

	logTrace.RegisterResponse(response)

	return response, nil
}

func (s Degree) scanRow(p pgx.Row) (model.Degree, error) {
	m := model.Degree{}

	specificLevelNull := sql.NullString{}

	err := p.Scan(
		&m.IdDegree,
		&m.NameDegree,
		&specificLevelNull,
	)
	if err != nil {
		return model.Degree{}, err
	}

	m.SpecificLevel = specificLevelNull.String

	return m, nil

}
