package started

import (
	"go-tel/src/telegram_bot/flow_started/schema/hello/keyboard"
	tele "gopkg.in/telebot.v3"
)

func SelectorNextHandler(c tele.Context) error {
	return c.Edit("hear edite help", keyboard.Selector)
}
