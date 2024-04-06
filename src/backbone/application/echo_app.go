package application

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-tel/src/backbone/api"
	echo_middleware "go-tel/src/backbone/api/middleware"
	_ "go-tel/src/backbone/docs"
	"go-tel/src/backbone/service_layer"
)

type App struct {
	app *echo.Echo
}

func (a *App) addMiddleware() {
	a.app.Use(echo_middleware.LoggerMiddleware)
}
func (a *App) initRoute() {
	apiRouter := service_layer.APIRouter{Router: a.app}
	api.InitRoute(apiRouter.Group("/api"))
	a.app.GET("/swagger/*", echoSwagger.WrapHandler)
}
func (a *App) Start() {
	// middleware
	a.addMiddleware()

	// rotes
	a.initRoute()

	// started app
	a.app.Logger.Fatal(a.app.Start(":8080"))
}

func NewApp() {
	e := echo.New()
	app := App{app: e}
	app.Start()

}
