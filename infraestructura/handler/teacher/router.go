package teacher

import (
	"notas/domain/teacher"
	storageTeacher "notas/infraestructura/postgres/teacher"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := teacher.New(storageTeacher.New(specification.DB))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routesTeacher := api.Group("teacher", middlewares...)

	routesTeacher.POST("", h.create)
	routesTeacher.GET("", h.getWhereAll)
	routesTeacher.GET("/getWhere", h.getWhere)
	routesTeacher.PATCH("/update", h.update)

}
