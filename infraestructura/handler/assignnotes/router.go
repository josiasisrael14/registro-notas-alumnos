package assignnotes

import (
	"notas/domain/assignnotes"
	storageAssignnotes "notas/infraestructura/postgres/assignnotes"
	"notas/model"

	"github.com/gin-gonic/gin"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := assignnotes.New(storageAssignnotes.New(specification.DB))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routesDegree := api.Group("assignnotes", middlewares...)

	routesDegree.GET("", h.getWhereAll)

}
