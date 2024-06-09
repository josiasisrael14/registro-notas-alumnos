package subjecTeacher

import (
	"context"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/jackc/pgx/v5"
)

var (
	_psqlGetAllWhere = ` SELECT 
                         ts.subjects_teacher_id, 
                         ts.comments, 
                         t.name,
						 t.surnames, 
                         g.grade_name, 
                         s.secction, 
                         m.nombremateria
                         FROM 
                         teacher_subjects ts
                         INNER JOIN 
                         teachers t ON ts.id_teacher = t.id_teacher
                         INNER JOIN 
                         grades g ON ts.grades_id = g.grades_id
                         INNER JOIN 
                         section s ON ts.section_id = s.section_id
                         INNER JOIN 
                         materias m ON ts.materiaid = m.materiaid`
)

var (
	_psqlGetWhere = ` SELECT 
                         ts.subjects_teacher_id, 
                         ts.comments, 
                         t.name,
						 t.surnames, 
                         g.grade_name, 
                         s.secction, 
                         m.nombremateria
                         FROM 
                         teacher_subjects ts
                         INNER JOIN 
                         teachers t ON ts.id_teacher = t.id_teacher
                         INNER JOIN 
                         grades g ON ts.grades_id = g.grades_id
                         INNER JOIN 
                         section s ON ts.section_id = s.section_id
                         INNER JOIN 
                         materias m ON ts.materiaid = m.materiaid`
)

var (
	_psqlInsertSubjectTeacher = `insert into teacher_subjects (id_teacher,grades_id,section_id,materiaid,comments)values($1,$2,$3,$4,$5);
  `
)

var (
	_psqlUpdateSubjectTeacher = `update teacher_subjects set id_teacher=$1,grades_id=$2,section_id,materiaid,comments where subjects_teacher_id=$3`
)

type SubjectTeacher struct {
	db model.PgxPool
}

func New(db model.PgxPool) SubjectTeacher {
	return SubjectTeacher{
		db: db,
	}
}

func (s SubjectTeacher) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.SubjectTeachers, error) {
	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAllWhere, specification)

	logTrace := register.NewPostgres(ctx, "postgres.subjectteacher.GetAllWhere")
	logTrace.RegisterRequest(query, args)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		logTrace.RegisterFailed(err)
		return nil, err
	}
	defer rows.Close()

	var ms model.SubjectTeachers
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

func (s SubjectTeacher) CreateSubjectTeacher(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error) {
	logTracer := register.NewPostgres(ctx, "postgres.subjecTeacher.CreateSubjectTeacher")

	logTracer.RegisterRequest(_psqlInsertSubjectTeacher, []any{request})

	_, err := s.db.Exec(ctx, _psqlInsertSubjectTeacher, request.NameTeacher, request.Grade, request.Section, request.Subject, request.Comments)
	if err != nil {
		logTracer.RegisterFailed(err)
		return model.ResponseStatusSubjectTeacher{}, err
	}
	response := model.ResponseStatusSubjectTeacher{
		Response: "subjectTeacher create",
	}

	logTracer.RegisterResponse(response)

	return response, nil
}

func (s SubjectTeacher) UpdateSubjectTeacher(ctx context.Context, request model.SubjectTeacher) (model.ResponseStatusSubjectTeacher, error) {
	logTrace := register.NewPostgres(ctx, "postgres.subjecTeacher.UpdateSubjectTeacher")

	logTrace.RegisterRequest(_psqlUpdateSubjectTeacher, []any{request})

	_, err := s.db.Exec(ctx, _psqlUpdateSubjectTeacher, request.NameTeacher, request.Grade, request.Section, request.Subject, request.Comments)

	if err != nil {
		logTrace.RegisterFailed(err)
		return model.ResponseStatusSubjectTeacher{}, err
	}

	response := model.ResponseStatusSubjectTeacher{

		Response: "subjectTeacher update",
	}

	logTrace.RegisterResponse(response)

	return response, nil
}

func (s SubjectTeacher) GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.SubjectTeacher, error) {
	query, args := repository.BuildQueryAndArgs(_psqlGetWhere, specification)

	logTracer := register.NewPostgres(ctx, "postgres.subjecTeacher.GetWhere")
	logTracer.RegisterRequest(query, args)

	m, err := s.scanRow(s.db.QueryRow(ctx, query, args...))
	if err != nil {
		logTracer.RegisterFailed(err)

		return model.SubjectTeacher{}, err
	}
	logTracer.RegisterResponse(m)

	return m, nil
}

func (t SubjectTeacher) scanRow(p pgx.Row) (model.SubjectTeacher, error) {
	m := model.SubjectTeacher{}

	err := p.Scan(
		&m.IdSubjectTeacher,
		&m.Comments,
		&m.NameTeacher,
		&m.Surnames,
		&m.Grade,
		&m.Section,
		&m.Subject,
	)
	if err != nil {
		return model.SubjectTeacher{}, err
	}

	return m, nil

}
