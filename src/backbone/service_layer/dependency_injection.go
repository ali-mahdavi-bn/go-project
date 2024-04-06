package service_layer

import (
	"reflect"
	"runtime"
	"strings"
)

func InjectHandlers(commands []FuncHandler) map[string]FuncHandler {
	injectedCommandHandlers := make(map[string]FuncHandler)
	for _, v := range commands {
		command := getNameCommandHandler(v)
		injectedCommandHandlers[command] = v
	}
	return injectedCommandHandlers
}

func injectDependency(Dependencies map[string]any, commandDependencies []string) map[string]any {
	commandAndCommandHandlers := make(map[string]any)
	for _, i := range commandDependencies {
		if v, ok := Dependencies[i]; ok {
			commandAndCommandHandlers[i] = v
		}
	}
	return commandAndCommandHandlers
}

func getNameCommandHandler(i interface{}) string {
	if funcName := getNameFunc(i); funcName != "" {
		return funcName[:len(funcName)-7]
	}
	return ""
}

func getNameFunc(i interface{}) string {
	fullFuncName := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")
	funcName := fullFuncName[1]
	if len(fullFuncName) == 2 {
		return funcName
	}
	return ""
}
