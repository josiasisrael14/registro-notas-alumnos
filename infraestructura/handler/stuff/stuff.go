package stuff

import (
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-repository-odl.git/repository"
	"github.com/gin-gonic/gin"

	"notas/domain/stuff"
	"notas/model"
)

type handler struct {
	useCase  stuff.UseCase
	response response.ApiResponse
}

func newHandler(useCase stuff.UseCase, response response.ApiResponse) handler {
	return handler{useCase: useCase, response: response}
}

func (h handler) create(c *gin.Context) {

	var stuff model.Stuff

	if err := c.BindJSON(&stuff); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	m, err := h.useCase.Create(c.Request.Context(), stuff)

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Create()", err))
		return
	}

	c.JSON(h.response.Created(c, m))

}

func (h handler) getWhere(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	idStuff := c.Query("idStuff")

	m, err := h.useCase.GetWhere(c.Request.Context(), idStuff)

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

	var stuff model.Stuff

	if err := c.BindJSON(&stuff); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	m, err := h.useCase.Update(c.Request.Context(), stuff)
	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Update()", err))
		return
	}
	c.JSON(h.response.Created(c, m))
}

func (h handler) delete(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	idStuff := c.Query("idStuff")

	_, err := h.useCase.Delete(c.Request.Context(), idStuff)
	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Delete()", err))
		return
	}

	c.JSON(h.response.Deleted(c))
}
