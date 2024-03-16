package keyboard

import (
	"go-tel/src/telegram_bot/flow_started/content"
	tele "gopkg.in/telebot.v3"
)

var (
	Menu        = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnHelp     = Menu.Text(content.BtnHelp)
	BtnSettings = Menu.Text(content.BtnSettings)
)
