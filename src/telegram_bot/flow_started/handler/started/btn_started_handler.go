package started

import (
	"go-tel/src/telegram_bot/flow_started/schema/hello"
	"go-tel/src/telegram_bot/flow_started/schema/hello/keyboard"
	tele "gopkg.in/telebot.v3"
)

func BtnHelpHandler(c tele.Context) error {
	hello.InitSchemaHelloSelector()
	return c.Send("kooomakkkkk", keyboard.Selector)
}
