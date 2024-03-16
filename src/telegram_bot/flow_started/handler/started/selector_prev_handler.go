package started

import (
	"go-tel/src/telegram_bot/flow_started/schema/hello"
	"go-tel/src/telegram_bot/flow_started/schema/hello/keyboard"
	tele "gopkg.in/telebot.v3"
)

func SelectorPrevHandler(c tele.Context) error {
	hello.InitSchemaBtnPrev()
	return c.Edit("hear edite help ascdascasdc", keyboard.Selector2)
}
