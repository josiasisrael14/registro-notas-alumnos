package section

import (
	"notas/domain/section"
	storageSection "notas/infraestructura/postgres/section"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := section.New(storageSection.New(specification.DB))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routesSection := api.Group("section", middlewares...)

	routesSection.POST("", h.create)
	routesSection.GET("", h.getWhereAll)
	routesSection.GET("/getWhere", h.getWhere)
	routesSection.PATCH("/update", h.update)
}
