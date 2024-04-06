package entrypoints

import (
	"go-tel/src/backbone/service_layer"
)

func InitProductsRoutes(e *service_layer.APIRouter) {
	e.Tag = "products"

	service_layer.Handle(e, "GET", "/get-two-user", GetTwoUser, &service_layer.SwaggerConfig{})
	service_layer.Handle(e, "GET", "/bbaa", GetTwoUser, &service_layer.SwaggerConfig{
		Summary: "Test One",
	})
}
