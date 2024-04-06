package service_layer

import (
	"github.com/labstack/echo/v4"
)

type IEventDispatcher interface {
	Dispatch(events []IBaseEvent)
}
type HandlerFunc echo.HandlerFunc
type IBaseEvent interface {
}

type Command interface {
	GetDependencies() []string
	SetDependencies(dependencies []string)
}

type CommandHandles struct {
	Request      echo.Context
	Dependencies []string
}

func (c *CommandHandles) GetDependencies() []string {

	return c.Dependencies
}
func (c *CommandHandles) SetDependencies(dependencies []string) {
	c.Dependencies = dependencies
}

type FuncHandler func(c echo.Context, command Command, dependencies map[string]any) error
