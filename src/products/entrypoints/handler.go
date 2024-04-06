package entrypoints

import (
	"github.com/labstack/echo/v4"
	"go-tel/src/products"
	"go-tel/src/products/service_layer/handlers"
)

var (
	bootstrap = products.Bootstrap()
)

func GetTwoUser(c echo.Context) error {
	command := handlers.NewHelloCommand("ali mah")
	return bootstrap.Handle(c, command)
}
