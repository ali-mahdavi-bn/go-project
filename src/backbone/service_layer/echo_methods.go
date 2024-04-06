package service_layer

import "github.com/labstack/echo/v4"

func MethodGetEcho(e *APIRouter, path string, handler HandlerFunc) {
	e.Router.GET(path, echo.HandlerFunc(handler))
}

func MethodPostEcho(e *APIRouter, path string, handler HandlerFunc) {
	e.Router.POST(path, echo.HandlerFunc(handler))
}

func MethodPutEcho(e *APIRouter, path string, handler HandlerFunc) {
	e.Router.PUT(path, echo.HandlerFunc(handler))
}
