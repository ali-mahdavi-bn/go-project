package application

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"time"
)

type Telegram struct {
	bot *tele.Bot
}

func NewTelegram() {
	pref := tele.Settings{
		Token:  "7066018155:AAGTr1TPb7N_bbj2ASxarrLzl85xVQ4sXww",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	telegram := Telegram{
		bot: b,
	}
	telegram.Start()
}

func (t *Telegram) addMiddleware() {
	t.bot.Use(middleware.Logger())
	t.bot.Use(middleware.AutoRespond())
}
func (t *Telegram) addHandlers() {
	LoadHandlers(t.bot)
}

func (t *Telegram) Start() {
	// middleware
	t.addMiddleware()

	// Load Handler
	t.addHandlers()

	// start telegram bot
	t.bot.Start()
}
