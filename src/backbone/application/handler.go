package application

import (
	"go-tel/src/telegram_bot/flow_started"
	tele "gopkg.in/telebot.v3"
)

func LoadHandler(bot *tele.Bot) {
	for command, handler := range flow_started.LoadHandlerHello() {
		bot.Handle(command, handler)
	}
}
