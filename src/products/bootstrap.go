package products

import (
	"go-tel/src"
	"go-tel/src/backbone/service_layer"
	"go-tel/src/products/service_layer/handlers"
)

func Bootstrap() service_layer.MessageBus {
	dependencies := map[string]any{
		"uow": src.NewUnitOfWork(),
	}

	commandHandlers := service_layer.InjectHandlers(handlers.Handlers)
	return service_layer.MessageBus{
		CommandHandlers: commandHandlers,
		Dependencies:    dependencies,
	}
}
