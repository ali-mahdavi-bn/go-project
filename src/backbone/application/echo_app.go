package application

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-tel/src/backbone/api"
	echo_middleware "go-tel/src/backbone/api/middleware"
	_ "go-tel/src/backbone/docs"
	"go-tel/src/backbone/service_layer"
)

func InitAppEcho() {
	e := echo.New()

	// middleware
	e.Use(echo_middleware.LoggerMiddleware)

	// rotes
	apiRouter := service_layer.APIRouter{Router: e}
	api.InitRoute(apiRouter.Group("/api"))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// started app
	e.Logger.Fatal(e.Start(":8080"))

}
