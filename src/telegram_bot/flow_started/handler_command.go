package flow_started

import (
	"go-tel/src/telegram_bot/flow_started/handler/started"
	"go-tel/src/telegram_bot/flow_started/schema/hello/keyboard"
	tele "gopkg.in/telebot.v3"
)

func LoadHandlerStarted() map[any]func(m tele.Context) error {
	commandMap := make(map[any]func(c tele.Context) error)

	// command
	commandMap["/start"] = started.StartedHandle

	// keyboard
	commandMap[&keyboard.SelectorPrev] = started.SelectorPrevHandler
	commandMap[&keyboard.SelectorNext] = started.SelectorNextHandler
	commandMap[&keyboard.BtnHelp] = started.BtnHelpHandler
	return commandMap
}
