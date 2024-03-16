package api

import (
	"go-tel/src/backbone/service_layer"
	products "go-tel/src/products/entrypoints"
)

func InitRoute(e *service_layer.APIRouter) {
	// Product
	products.InitProductsRoutes(e)
}
