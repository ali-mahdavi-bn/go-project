package domain

import "go-tel/src/backbone/service_layer"

type HelloCommand struct {
	service_layer.CommandHandles
	Ali string
}

type Hello2Command struct {
	service_layer.CommandHandles
}
