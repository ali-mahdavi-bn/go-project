package hello

import (
	"go-tel/src/telegram_bot/flow_started/schema/hello/keyboard"
	tele "gopkg.in/telebot.v3"
)

func InitSchemaHelloSelector() *tele.ReplyMarkup {
	keyboard.Selector.Inline(
		keyboard.Selector.Row(keyboard.SelectorPrev),
	)
	return keyboard.Selector
}

func InitSchemaBtnPrev() *tele.ReplyMarkup {
	keyboard.Selector2.Inline(
		keyboard.Selector2.Row(keyboard.SelectorNext),
	)
	return keyboard.Selector2
}

func InitSchemaHelloBtn() *tele.ReplyMarkup {
	keyboard.Menu.Reply(
		keyboard.Menu.Row(keyboard.BtnHelp, keyboard.BtnSettings),
	)
	return keyboard.Menu
}
