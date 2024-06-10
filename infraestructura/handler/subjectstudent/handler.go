package subjectstudent

import (
	"net/http"
	"notas/domain/subjectstudent"
	"notas/model"
	"notas/response"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-domain-alicorp.git/customerror"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase  subjectstudent.UseCase
	response response.ApiResponse1
}

func newHandler(useCase subjectstudent.UseCase, response response.ApiResponse1) handler {
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

	var studentSubject model.StudentSubject

	if err := c.BindJSON(&studentSubject); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	m, err := h.useCase.GetAllWhere(c.Request.Context(), repository.FieldsSpecification{})

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.GetAllWhere()", err))
		return
	}

	for _, existingStudentSubject := range m {
		if existingStudentSubject.IdStudent == studentSubject.NameStudent {
			customErr := customerror.NewError()
			customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
			customErr.Fields.AddGeneric("requestBackend", "Student is already enrolled", customerror.IssueRequestBackendFailed)
			c.JSON(http.StatusUnprocessableEntity, customErr)
			return
		}
	}

	rs, err := h.useCase.CreateSubjectStudent(c.Request.Context(), studentSubject)

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Create()", err))
		return
	}

	c.JSON(h.response.Created(c, rs))

}
