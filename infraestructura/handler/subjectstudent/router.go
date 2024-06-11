package subjectstudent

import (
	"notas/domain/subjectstudent"
	StorageSubjecStudent "notas/infraestructura/postgres/subjectstudent"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"

	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := subjectstudent.New(StorageSubjecStudent.New(specification.DB))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routesSubjectStudent := api.Group("subjectStudent", middlewares...)

	routesSubjectStudent.GET("", h.getWhereAll)
	routesSubjectStudent.POST("", h.create)
	routesSubjectStudent.GET("/id", h.getWhere)
}
