package student

import (
	"notas/domain/student"
	"notas/model"
	"time"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/gin-gonic/gin"
)

type StudentTemp struct {
	StudentId   string `json:"studentId"`
	NameStudent string `json:"nameStudent"`
	LastName    string `json:"lastName"`
	BirthDate   string `json:"birthDate"`
}

type handler struct {
	useCase  student.UseCase
	response response.ApiResponse
}

func newHandler(useCase student.UseCase, response response.ApiResponse) handler {
	return handler{useCase: useCase, response: response}
}

func (h handler) create(c *gin.Context) {

	var student StudentTemp

	if err := c.BindJSON(&student); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	birthDate, err := time.Parse(model.DateFormat, student.BirthDate)
	//birthDate, err := dateparser.ParseDate(student.BirthDate)

	if err != nil {
		c.JSON(h.response.Error(c, "Invalid date format", err))
		return
	}

	studentRequest := model.Student{
		NameStudent: student.NameStudent,
		LastName:    student.LastName,
		BirthDate:   birthDate,
	}

	m, err := h.useCase.Create(c.Request.Context(), studentRequest)
	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Create()", err))
		return
	}

	c.JSON(h.response.Created(c, m))

}

func (h handler) getWhere(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	studentId := c.Query("studentId")

	m, err := h.useCase.GetWhere(c.Request.Context(), studentId)

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.GetWhere()", err))
		return
	}

	c.JSON(h.response.OK(c, m))
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

func (h handler) update(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	var student model.Student

	if err := c.BindJSON(&student); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	m, err := h.useCase.Update(c.Request.Context(), student)
	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Update()", err))
		return
	}
	c.JSON(h.response.Created(c, m))
}

func (h handler) delete(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	studentId := c.Query("studentId")

	_, err := h.useCase.Delete(c.Request.Context(), studentId)
	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Delete()", err))
		return
	}

	c.JSON(h.response.Deleted(c))
}
