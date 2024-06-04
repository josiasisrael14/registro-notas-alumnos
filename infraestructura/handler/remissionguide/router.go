package remissionguide

import (
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"github.com/gin-gonic/gin"

	"notas/domain/remissionguide"
	OrderSoap "notas/infraestructura/soap/order"
	"notas/model"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := remissionguide.New(OrderSoap.New(specification.ConfigSoaps))

	return newHandler(useCase, response.New(response.NewDefaltResponse()))
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/odoo", middlewares...)

	routes.GET("order", h.retrieveSupplyReception)
}
