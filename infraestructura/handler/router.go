package handler

import (
	"notas/infraestructura/handler/assignnotes"
	"notas/infraestructura/handler/degree"
	"notas/infraestructura/handler/remissionguide"
	"notas/infraestructura/handler/section"
	"notas/infraestructura/handler/student"
	"notas/infraestructura/handler/stuff"
	subjectTeacher "notas/infraestructura/handler/subjectTeacher"
	"notas/infraestructura/handler/subjectstudent"
	"notas/infraestructura/handler/teacher"
	"notas/model"
)

func InitRoutes(specification model.RouterSpecification) {
	stuff.NewRouter(specification)
	degree.NewRouter(specification)
	section.NewRouter(specification)
	teacher.NewRouter(specification)
	student.NewRouter(specification)
	subjectTeacher.NewRouter(specification)
	remissionguide.NewRouter(specification)
	subjectstudent.NewRouter(specification)
	assignnotes.NewRouter(specification)
}
