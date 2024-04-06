package service_layer

import "github.com/labstack/echo/v4"

type APIRouter struct {
	Router             *echo.Echo
	Tag                string
	Path               string
	InitializerSwagger bool
}

func NewAPIRouter(InitializerSwagger bool) *APIRouter {
	return &APIRouter{InitializerSwagger: InitializerSwagger}
}
func (u *APIRouter) Group(prefix string) *APIRouter {
	u.Path += prefix
	return u
}

type SwaggerConfig struct {
	Path        string
	Method      string
	Tag         string
	Description string
	Summary     string
	Status      string
}
