package subjectTeacher

import (
	"notas/domain/subjectTeacher"
	StorageSubjecTeacher "notas/infraestructura/postgres/subjectTeacher"
	"notas/model"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := subjectTeacher.New(StorageSubjecTeacher.New(specification.DB))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routesSubjectTeacher := api.Group("subjectTeacher", middlewares...)

	routesSubjectTeacher.GET("", h.getWhereAll)
	routesSubjectTeacher.GET("/id", h.getWhere)
	routesSubjectTeacher.POST("", h.create)
	routesSubjectTeacher.PATCH("/update", h.update)

}
