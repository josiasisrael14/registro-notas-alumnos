package bootstrap

import (
	"context"
	"os"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/v2/apiutils/bootstrap"
	"github.com/joho/godotenv"

	//"github.com/rs/cors"
	"github.com/gin-contrib/cors"

	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"

	"notas/infraestructura/handler"
	"notas/model"
)

func Run2(boot []byte) {
	// we ignore the error because it loads the .env file, and in production the envs are defined differently
	_ = godotenv.Load()
	ctx := context.Background()
	applicationName := getApplicationName()

	db := newDatabase(ctx)
	ginEntry := newGinEntry(boot)
	ginEntry.Bootstrap(ctx)

	logger := newLogger()

	traceProvider := initTracer(ctx, applicationName)
	defer shutdownTraceProvider(ctx, logger, traceProvider)

	api := ginEntry.Router
	//apiGroup := api.Group("") // this group is needed for middlewares, omit paths on common rk
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Allow all origins (you can specify specific origins)
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}
	api.Use(cors.New(config))
	// Implementar CORS
	// Configurar el middleware CORS
	/*corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Agrega la URL de tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}*/
	//apiGroup.Use(cors.New(corsConfig))
	//apiGroup.Use(otelgin.Middleware(applicationName, otelgin.WithTracerProvider(traceProvider)))
	//apiGroup.Use(newLogTracerMiddleware(logger))
	//apiGroup.Use(newCommonMiddleware())
	//apiGroup.Use(corsMiddleware())
	//apiGroup.Use(middleware.RegisterHeadersDefaultResponse())

	handler.InitRoutes(model.RouterSpecification{
		Api:           api,
		DB:            db,
		Logger:        logger,
		ConfigService: model.ConfigService{},
		ConfigSoaps: model.ConfigSoaps{
			Url:      os.Getenv("URLGUIAEREMISION"),
			User:     os.Getenv("USERSOAP"),
			Password: os.Getenv("PASSWORD"),
		},
	})

	bootstrap.PrintRoutes(api, logger)

	rkentry.GlobalAppCtx.WaitForShutdownSig()
	ginEntry.Interrupt(ctx)
}
