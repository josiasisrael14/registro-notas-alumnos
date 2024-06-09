package subjectstudent

import (
	"context"
	"database/sql"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/jackc/pgx/v5"
)

var (
	_psqlGetAllWhere = ` SELECT 
                         ts.subject_student_id, 
                         st.name,
						 st.surnames, 
                         g.grade_name, 
						 g.specific_level,
                         s.secction 
                         FROM 
                         students_subjects ts
                         INNER JOIN 
                         students st ON ts.student_id = st.student_id
                         INNER JOIN 
                         grades g ON ts.grades_id = g.grades_id
                         INNER JOIN 
                         section s ON ts.section_id = s.section_id
                         `
)

type SubjectStudent struct {
	db model.PgxPool
}

func New(db model.PgxPool) SubjectStudent {
	return SubjectStudent{
		db: db,
	}
}

func (s SubjectStudent) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.StudentSubjects, error) {
	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAllWhere, specification)

	logTrace := register.NewPostgres(ctx, "postgres.subjectstudent.GetAllWhere")
	logTrace.RegisterRequest(query, args)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		logTrace.RegisterFailed(err)
		return nil, err
	}
	defer rows.Close()

	var ms model.StudentSubjects
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

func (t SubjectStudent) scanRow(p pgx.Row) (model.StudentSubject, error) {
	m := model.StudentSubject{}

	specificationLevelNull := sql.NullString{}

	err := p.Scan(
		&m.IdSubjectStudent,
		&m.NameStudent,
		&m.LastName,
		&m.Grade,
		&specificationLevelNull,
		&m.Section,
	)
	if err != nil {
		return model.StudentSubject{}, err
	}

	m.SpecificationLevel = specificationLevelNull.String

	return m, nil

}
