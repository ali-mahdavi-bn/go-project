package keyboard

import (
	"go-tel/src/telegram_bot/flow_started/content"
	tele "gopkg.in/telebot.v3"
)

var (
	Selector  = &tele.ReplyMarkup{}
	Selector2 = &tele.ReplyMarkup{}

	SelectorPrev = Selector.Data(content.SelectorPrev, content.SelectorPrevText)
	SelectorNext = Selector2.Data(content.SelectorNext, content.SelectorNextText)
)
