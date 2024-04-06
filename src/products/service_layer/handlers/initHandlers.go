package handlers

import "go-tel/src/backbone/service_layer"

var Handlers = []service_layer.FuncHandler{
	HelloHandler,
}
