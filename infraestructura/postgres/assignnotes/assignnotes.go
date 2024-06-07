package assignnotes

import (
	"context"
	"database/sql"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/jackc/pgx/v5"
)

var (
	_psqlGetAllWhere = ` select n.note_id, 
	                            s.name,
								t.name,
								m.nombremateria,
								g.grade_name,
								g.specific_level,
								se.secction,
								n.note,
								n.date,
								n.comments
								from notas n
								 inner join students s
								on n.student_id=s.student_id
	inner join teachers t on n.id_teacher=t.id_teacher inner join materias m on n.materiaid=m.materiaid inner join grades g on n.grades_id=g.grades_id 
	inner join section se on n.section_id=se.section_id
                         `
)

type AssignNotes struct {
	db model.PgxPool
}

func New(db model.PgxPool) AssignNotes {
	return AssignNotes{
		db: db,
	}
}

func (a AssignNotes) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.AssignNotes, error) {
	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAllWhere, specification)

	logTrace := register.NewPostgres(ctx, "postgres.assignnotes.GetAllWhere")
	logTrace.RegisterRequest(query, args)

	rows, err := a.db.Query(ctx, query, args...)
	if err != nil {
		logTrace.RegisterFailed(err)
		return nil, err
	}
	defer rows.Close()

	var ms model.AssignNotes
	for rows.Next() {
		m, err := a.scanRow(rows)
		if err != nil {
			logTrace.RegisterFailed(err)
			return nil, err
		}
		ms = append(ms, m)
	}

	logTrace.RegisterResponse(ms)

	return ms, nil
}

func (t AssignNotes) scanRow(p pgx.Row) (model.AssignNote, error) {
	m := model.AssignNote{}

	specificationLevelNull := sql.NullString{}

	err := p.Scan(
		&m.Note_id,
		&m.NameStudent,
		&m.NameTeacher,
		&m.NameStuff,
		&m.Grade,
		&specificationLevelNull,
		&m.Section,
		&m.Note,
		&m.Date,
		&m.Comments,
	)
	if err != nil {
		return model.AssignNote{}, err
	}

	m.SpecificationLevel = specificationLevelNull.String

	return m, nil

}
