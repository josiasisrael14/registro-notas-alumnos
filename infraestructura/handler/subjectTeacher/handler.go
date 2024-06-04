package subjectTeacher

import (
	subjectteacher "notas/domain/subjectTeacher"
	"notas/model"
	"notas/response"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase  subjectteacher.UseCase
	response response.ApiResponse1
}

func newHandler(useCase subjectteacher.UseCase, response response.ApiResponse1) handler {
	return handler{useCase: useCase, response: response}
}

func (h handler) getWhereAll(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	m, err := h.useCase.GetAllWhere(c.Request.Context(), repository.FieldsSpecification{})

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.GetAllWhere()", err))
		return
	}

	c.JSON(h.response.OK(c, m))
}

func (h handler) create(c *gin.Context) {

	var subjectTeacher model.SubjectTeacher

	if err := c.BindJSON(&subjectTeacher); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	m, err := h.useCase.CreateSubjectTeacher(c.Request.Context(), subjectTeacher)

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.CreateSubjectTeacher()", err))
		return
	}

	c.JSON(h.response.Created(c, m))

}

func (h handler) update(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	var subjectTeacher model.SubjectTeacher

	if err := c.BindJSON(&subjectTeacher); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	m, err := h.useCase.Update(c.Request.Context(), subjectTeacher)
	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Update()", err))
		return
	}
	c.JSON(h.response.Created(c, m))
}

func (h handler) getWhere(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	idSubjectTeacher := c.Query("idSubjectTeacher")

	m, err := h.useCase.GetWhere(c.Request.Context(), idSubjectTeacher)

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.GetWhere()", err))
		return
	}

	c.JSON(h.response.OK(c, m))
}
