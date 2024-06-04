package remissionguide

import (
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/request"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/response"
	"github.com/gin-gonic/gin"

	"notas/domain/remissionguide"
	"notas/model"
)

type handler struct {
	useCase  remissionguide.UseCase
	response response.ApiResponse
}

func newHandler(useCase remissionguide.UseCase, response response.ApiResponse) handler {
	return handler{useCase: useCase, response: response}
}

func (h handler) retrieveSupplyReception(c *gin.Context) {
	reasonConsultation := c.Query("reasonConsultation")
	referenceType := c.Query("referenceType")
	society := c.Query("society")
	supplierCode := c.Query("supplierCode")
	supplyCenter := c.Query("supplyCenter")
	documentType := c.Query("documentType")
	dateStart := c.Query("dateStart")
	dateEnd := c.Query("dateEnd")
	documentNumber := c.Query("documentNumber")
	documentPosition := c.Query("documentPosition")
	materialNumber := c.Query("materialNumber")
	destinationCenter := c.Query("destinationCenter")

	fieldsForValidate := map[string]string{
		"reasonConsultation": reasonConsultation,
		"referenceType":      referenceType,
		"society":            society,
		"supplierCode":       supplierCode,
		"dateStart":          dateStart,
		"dateEnd":            dateEnd,
	}

	if err := request.ValidateMissingParam(fieldsForValidate); err != nil {
		c.JSON(h.response.ParamFailed(c, err))
		return
	}

	ms, err := h.useCase.GetOrder(c.Request.Context(), model.RemissionOrderRequest{
		ReasonConsultation: reasonConsultation,
		ReferenceType:      referenceType,
		Society:            society,
		SupplierCode:       supplierCode,
		SupplyCenter:       supplyCenter,
		DocumentType:       documentType,
		DateStart:          dateStart,
		DateEnd:            dateEnd,
		DocumentNumber:     documentNumber,
		DocumentPosition:   documentPosition,
		MaterialNumber:     materialNumber,
		DestinationCenter:  destinationCenter,
	})

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.GetOrder()", err))
		return
	}

	c.JSON(h.response.OK(c, ms))
}
