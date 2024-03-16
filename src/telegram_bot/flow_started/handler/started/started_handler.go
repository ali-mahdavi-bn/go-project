package started

import (
	"go-tel/src/telegram_bot/flow_started/content"
	"go-tel/src/telegram_bot/flow_started/schema/hello"
	tele "gopkg.in/telebot.v3"
)

func StartedHandle(c tele.Context) error {
	keyboard := hello.InitSchemaHelloBtn()
	return c.Send(content.FirstMessage, keyboard)
}
