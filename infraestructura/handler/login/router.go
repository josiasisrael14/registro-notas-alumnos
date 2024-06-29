package login

import (
	loginStorage "notas/infraestructura/postgres/login"
	"notas/model"

	"notas/domain/login"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := login.New(loginStorage.New(specification.DB))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routesSection := api.Group("login", middlewares...)

	routesSection.POST("", h.authenticate)
}
