package products

import (
	"go-tel/src"
	"go-tel/src/backbone/service_layer"
)

var Dependencies = map[string]any{
	"uow": src.NewGormUnitOfWork(),
}
var name = src.NewGormUnitOfWork()

type Bootstrap struct {
	commands map[string]service_layer.Command
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		commands: make(map[string]service_layer.Command),
	}
}

func (h *Bootstrap) Handle(command service_layer.Command) error {
	commandDependencies := command.GetDependencies()
	injectDependency := make(map[string]any)
	for _, i := range commandDependencies {
		if v, ok := Dependencies[i]; ok {
			injectDependency[i] = v
		}
	}
	resultCommandHandler := command.CommandHandler(injectDependency)
	return resultCommandHandler
}
