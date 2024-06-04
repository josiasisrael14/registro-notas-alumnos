package stuff

import (
	"context"
	"database/sql"
	"net/http"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-domain-alicorp.git/customerror"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/jackc/pgx/v5"

	"notas/model"
)

var _fieldWhere = []string{
	"materiaid",
	"nombremateria",
	"descripcion",
}

const _table = "materias"

var (
	_psqlGetWhere    = repository.BuildSQLSelectFields(_table, _fieldWhere)
	_psqlGetAllWhere = repository.BuildSQLSelectFields(_table, _fieldWhere)
)

var (
	_psqlInsertMateria = `INSERT INTO materias( nombremateria,descripcion )VALUES($1,$2)`
)

var (
	_psqlUpdateStuff = `update materias set nombremateria=$1,descripcion=$2 where materiaid=$3`
)

var (
	_psqlDeleteStuff = `delete from materias where materiaid=$1`
)

type Stuff struct {
	db model.PgxPool
}

func New(db model.PgxPool) Stuff {
	return Stuff{
		db: db,
	}
}

func (s Stuff) CreateStuff(ctx context.Context, request model.Stuff) (model.ResponseStatusStuff, error) {
	logTracer := register.NewPostgres(ctx, "postgres.stuff.CreateStuff")

	logTracer.RegisterRequest(_psqlInsertMateria, []any{request})

	if err := validateResponse(request); err != nil {
		return model.ResponseStatusStuff{}, err
	}

	_, err := s.db.Exec(ctx, _psqlInsertMateria, request.NameStuff, request.Description)
	if err != nil {
		logTracer.RegisterFailed(err)
		return model.ResponseStatusStuff{}, err
	}
	response := model.ResponseStatusStuff{
		Response: "Stuff create",
	}

	logTracer.RegisterResponse(response)

	return response, nil
}

func (s Stuff) GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Stuff, error) {
	query, args := repository.BuildQueryAndArgs(_psqlGetWhere, specification)

	logTracer := register.NewPostgres(ctx, "postgres.stuff.GetWhere")
	logTracer.RegisterRequest(query, args)

	m, err := s.scanRow(s.db.QueryRow(ctx, query, args...))
	if err != nil {
		logTracer.RegisterFailed(err)

		return model.Stuff{}, err
	}
	logTracer.RegisterResponse(m)

	return m, nil
}

func (s Stuff) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Stuffs, error) {
	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAllWhere, specification)

	logTrace := register.NewPostgres(ctx, "postgres.stuff.GetAllWhere")
	logTrace.RegisterRequest(query, args)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		logTrace.RegisterFailed(err)
		return nil, err
	}
	defer rows.Close()

	var ms model.Stuffs
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

func (s Stuff) UpdateStuff(ctx context.Context, request model.Stuff) (model.ResponseStatusStuff, error) {
	logTrace := register.NewPostgres(ctx, "postgres.stuff.UpdateStuff")

	logTrace.RegisterRequest(_psqlUpdateStuff, []any{request})

	_, err := s.db.Exec(ctx, _psqlUpdateStuff, request.NameStuff, request.Description, request.IdStuff)

	if err != nil {
		logTrace.RegisterFailed(err)
		return model.ResponseStatusStuff{}, err
	}

	response := model.ResponseStatusStuff{

		Response: "update success",
	}

	logTrace.RegisterResponse(response)

	return response, nil
}

func (s Stuff) DeleteStuff(ctx context.Context, id string) (model.ResponseStatusStuff, error) {
	logTrace := register.NewPostgres(ctx, "postgres.stuff.DeleteStuff")

	logTrace.RegisterRequest(_psqlDeleteStuff, []any{id})

	_, err := s.db.Exec(ctx, _psqlDeleteStuff, id)

	if err != nil {
		logTrace.RegisterFailed(err)
		return model.ResponseStatusStuff{}, err
	}

	response := model.ResponseStatusStuff{

		Response: "delete success",
	}

	logTrace.RegisterResponse(response)

	return response, nil

}

func (s Stuff) scanRow(p pgx.Row) (model.Stuff, error) {
	m := model.Stuff{}

	nameStuffNull := sql.NullString{}
	descriptionNull := sql.NullString{}

	err := p.Scan(
		&m.IdStuff,
		&nameStuffNull,
		&descriptionNull,
	)
	if err != nil {
		return model.Stuff{}, err
	}

	m.NameStuff = nameStuffNull.String
	m.Description = descriptionNull.String

	return m, nil

}

func validateResponse(request model.Stuff) error {
	customErr := customerror.NewError()
	if request.NameStuff == "" {
		customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
		customErr.Fields.AddGeneric("requestBackend", "nameStuff can't be empty", customerror.IssueRequestBackendFailed)
		return customErr
	}

	if request.Description == "" {
		customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
		customErr.Fields.AddGeneric("requestBackend", " description can't be empty ", customerror.IssueRequestBackendFailed)
		return customErr
	}

	return nil
}
