package login

import (
	"context"
	"fmt"
	"net/http"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-domain-alicorp.git/customerror"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/jackc/pgx/v5"
)

/*var _fieldWhere = []string{
	"id_teacher",
	"name",
	"surnames",
}

const _table = "teachers"

var (
	_psqlGetWhere = repository.BuildSQLSelectFields(_table, _fieldWhere)
)*/

var (
	_psqlGetWhereTeacher = `select 
	                        t.id_teacher,
							t.name,
							t.surnames,
							ur.roles_id
							from teachers t
						    inner join user_role ur
							 on t.roles_id =ur.roles_id 
                         `
)

const Users = "1"
const Administrador = "2"

type Login struct {
	db model.PgxPool
}

func New(db model.PgxPool) Login {
	return Login{
		db: db,
	}
}

func (l Login) LoginAcceso(ctx context.Context, specification repository.FieldsSpecification) (model.ResponseStatusLogin, error) {
	query, args := repository.BuildQueryAndArgs(_psqlGetWhereTeacher, specification)

	logTracer := register.NewPostgres(ctx, "postgres.login.LoginAcceso")
	logTracer.RegisterRequest(query, args)

	m, _ := l.scanRow(l.db.QueryRow(ctx, query, args...))
	/*if err != nil {
		logTracer.RegisterFailed(err)
		return model.Login{}, err
	}*/

	if err := validateResponse(m); err != nil {
		return model.ResponseStatusLogin{}, fmt.Errorf("postgres.login.LoginAcceso().validateResponse(): %w", err)
	}

	if m.RolId == Users {

		m.Response = "usuario"
	} else {
		m.Response = "administrador"
	}

	logTracer.RegisterResponse(m)

	return m, nil
}

func (l Login) scanRow(p pgx.Row) (model.ResponseStatusLogin, error) {
	m := model.ResponseStatusLogin{}

	err := p.Scan(
		&m.IdTeacher,
		&m.NameTeacher,
		&m.Surname,
		&m.RolId,
	)
	if err != nil {
		return model.ResponseStatusLogin{}, err
	}

	return m, nil

}

func validateResponse(m model.ResponseStatusLogin) error {
	if m.IdTeacher == "" && m.NameTeacher == "" && m.Surname == "" {
		customErr := customerror.NewError()
		customErr.SetStatusHTTP(http.StatusUnauthorized)
		customErr.Fields.Add(customerror.ErrorDetail{
			Field:       "body",
			Issue:       customerror.IssueBodyBackendError,
			Description: "credencials invalidas no autorizado",
		})

		return customErr
	}

	return nil
}
