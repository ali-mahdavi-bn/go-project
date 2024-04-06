package service_layer

import (
	"strings"
)

func Handle(e *APIRouter, method string, path string, handler HandlerFunc, swaggerConfig *SwaggerConfig) {
	fullPath := e.Path + path

	if e.InitializerSwagger != true {
		HandleMethod(e, method, fullPath, handler)
	} else {
		setupSwaggerConfig(e, fullPath, method, handler, swaggerConfig)
		GenerateSwaggerConfig(swaggerConfig)
	}

}

func HandleMethod(e *APIRouter, method string, path string, handler HandlerFunc) {
	switch strings.ToLower(method) {
	case "get":
		MethodGetEcho(e, path, handler)
	case "post":
		MethodPostEcho(e, path, handler)
	case "put":
		MethodPutEcho(e, path, handler)
	}
}
