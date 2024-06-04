package student

import (
	"notas/domain/student"
	storageStudent "notas/infraestructura/postgres/student"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := student.New(storageStudent.New(specification.DB))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routesStudent := api.Group("student", middlewares...)

	routesStudent.POST("", h.create)
	routesStudent.GET("", h.getWhere)
	routesStudent.PATCH("/update", h.update)
	routesStudent.DELETE("", h.delete)
	routesStudent.GET("/all", h.getWhereAll)

}
