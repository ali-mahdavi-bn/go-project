package service_layer

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-tel/src/backbone/helpers/logger"
	"net/http"
	"reflect"
	"strings"
)

type MessageBus struct {
	CommandHandlers map[string]FuncHandler
	Dependencies    map[string]any
}

func (m MessageBus) Handle(c echo.Context, command Command) error {
	var err error
	defer errorHandler(err, c)
	commandDependencies := command.GetDependencies()
	injectDependencyi := injectDependency(m.Dependencies, commandDependencies)
	commandName := GetNameCommand(command)
	//commandName := "Hello"
	handler, ok := m.CommandHandlers[commandName]
	if ok {
		resultCommandHandler := handler(c, command, injectDependencyi)
		if err != nil {
			return err
		}
		return resultCommandHandler
	}
	return fmt.Errorf("Can`t handling command %sCommand", commandName)
}

func GetNameCommand(command Command) string {
	val := reflect.ValueOf(command)
	commandFullName := strings.Split(val.String(), ".")
	commandName := commandFullName[1]

	if len(commandFullName) == 2 {
		return commandName[:len(commandName)-14]
	}
	return ""
}

func errorHandler(err error, request echo.Context) {
	if r := recover(); r != nil {
		logger.Error(r)
		err = request.JSON(http.StatusInternalServerError, "مشکلی پیش آمده! لطفا با پشتیبان تماس بگیرین")
	}
}
