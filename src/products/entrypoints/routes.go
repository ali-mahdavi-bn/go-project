package entrypoints

import (
	"github.com/labstack/echo/v4"
	"go-tel/src/backbone/service_layer"
	"go-tel/src/products"
	"go-tel/src/products/service_layer/handlers"
)

func InitProductsRoutes(e *service_layer.APIRouter) {
	e.Tag = "products"
	bootstrap := products.NewBootstrap()

	service_layer.Handlers(e, "GET", "/aabbb",
		func(c echo.Context) error {
			command := handlers.NewHelloCommand(c)
			return bootstrap.Handle(command)
		}, map[string]any{
			"summary": "create user",
		})

	service_layer.Handlers(e, "GET", "/bbaa",
		func(c echo.Context) error {
			command := handlers.NewHelloCommand(c)
			return bootstrap.Handle(command)
		}, map[string]any{
			"summary": "create user",
		})
}
