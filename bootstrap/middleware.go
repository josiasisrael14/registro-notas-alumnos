package bootstrap

import (
	"notas/model"

	commonmiddleware "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/middleware"
	logtracemiddleware "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/middleware"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/v2/register"
	"github.com/gin-gonic/gin"
)

func newLogTracerMiddleware(logger model.Logger) gin.HandlerFunc {
	register.SetMaxCharResponse(getMaxCharBodyLogger())

	return logtracemiddleware.LoggerTracesHandler(logger, getDebugMode(), model.ErrorApi{
		Code:    model.InternalServerErrorCode,
		Message: model.InternalServerErrorDescription,
	})
}

func newCommonMiddleware() gin.HandlerFunc {
	return commonmiddleware.CommonMiddleware(model.ErrorApi{
		Code:    model.InternalServerErrorCode,
		Message: model.InternalServerErrorDescription,
	})
}

/*func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Configurar los encabezados CORS
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		// Verificar si es una solicitud OPTIONS (preflight)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Llamar al siguiente middleware
		c.Next()
	}
}*/
