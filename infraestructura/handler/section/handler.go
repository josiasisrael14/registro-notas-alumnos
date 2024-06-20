package section

import (
	"notas/domain/section"
	"notas/model"
	"notas/response"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase  section.UseCase
	response response.ApiResponse1
}

func newHandler(useCase section.UseCase, response response.ApiResponse1) handler {
	return handler{useCase: useCase, response: response}
}

func (h handler) create(c *gin.Context) {

	var section model.Section

	if err := c.BindJSON(&section); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	m, err := h.useCase.Create(c.Request.Context(), section)

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Create()", err))
		return
	}

	c.JSON(h.response.Created(c, m))

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

func (h handler) getWhere(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	idSection := c.Query("idSection")

	m, err := h.useCase.GetWhere(c.Request.Context(), idSection)

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.GetWhere()", err))
		return
	}

	c.JSON(h.response.OK(c, m))
}

func (h handler) update(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	var section model.Section

	if err := c.BindJSON(&section); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	m, err := h.useCase.Update(c.Request.Context(), section)
	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Update()", err))
		return
	}
	c.JSON(h.response.Created(c, m))
}
