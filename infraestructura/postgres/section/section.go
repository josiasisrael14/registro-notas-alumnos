package section

import (
	"context"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/jackc/pgx/v5"
)

var _fieldWhere = []string{
	"section_id",
	"secction",
}

const _table = "section"

var (
	//_psqlGetWhere    = repository.BuildSQLSelectFields(_table, _fieldWhere)
	_psqlGetAllWhere = repository.BuildSQLSelectFields(_table, _fieldWhere)
)

var (
	_psqlInsertSection = `INSERT INTO section(secction)VALUES($1)`
)

type Section struct {
	db model.PgxPool
}

func New(db model.PgxPool) Section {
	return Section{
		db: db,
	}
}

func (s Section) CreateSection(ctx context.Context, request model.Section) (model.ResponseStatusSection, error) {
	logTracer := register.NewPostgres(ctx, "postgres.section.CreateSection")

	logTracer.RegisterRequest(_psqlInsertSection, []any{request})

	_, err := s.db.Exec(ctx, _psqlInsertSection, request.NameSection)
	if err != nil {
		logTracer.RegisterFailed(err)
		return model.ResponseStatusSection{}, err
	}
	response := model.ResponseStatusSection{
		Response: "section create",
	}

	logTracer.RegisterResponse(response)

	return response, nil
}

func (s Section) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Sections, error) {
	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAllWhere, specification)

	logTrace := register.NewPostgres(ctx, "postgres.section.GetAllWhere")
	logTrace.RegisterRequest(query, args)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		logTrace.RegisterFailed(err)
		return nil, err
	}
	defer rows.Close()

	var ms model.Sections
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

func (s Section) scanRow(p pgx.Row) (model.Section, error) {
	m := model.Section{}

	err := p.Scan(
		&m.IdSection,
		&m.NameSection,
	)
	if err != nil {
		return model.Section{}, err
	}

	return m, nil

}
