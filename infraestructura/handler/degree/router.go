package degree

import (
	"notas/domain/degree"
	storageDegree "notas/infraestructura/postgres/degree"
	"notas/model"

	"github.com/gin-gonic/gin"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := degree.New(storageDegree.New(specification.DB))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routesDegree := api.Group("degree", middlewares...)

	routesDegree.POST("", h.create)
	routesDegree.GET("", h.getWhereAll)
	routesDegree.GET("/getWhere", h.getWhere)
	routesDegree.PATCH("/update", h.update)

}
