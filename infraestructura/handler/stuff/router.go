package stuff

import (
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"github.com/gin-gonic/gin"

	"notas/domain/stuff"
	storageStuff "notas/infraestructura/postgres/stuff"
	"notas/model"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := stuff.New(storageStuff.New(specification.DB))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routesMaterias := api.Group("stuff", middlewares...)

	routesMaterias.POST("", h.create)
	routesMaterias.GET("", h.getWhere)
	routesMaterias.PATCH("/update", h.update)
	routesMaterias.DELETE("", h.delete)
	routesMaterias.GET("/all", h.getWhereAll)

}
