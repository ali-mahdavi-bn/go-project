package application

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"time"
)

func TelegramBot() *tele.Bot {
	pref := tele.Settings{
		Token:  "7066018155:AAGTr1TPb7N_bbj2ASxarrLzl85xVQ4sXww",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return b
}

func InitTelegram(bot *tele.Bot) {

	// middleware
	bot.Use(middleware.Logger())
	bot.Use(middleware.AutoRespond())

	// Load Handler
	LoadHandler(bot)

	// start app
	bot.Start()
}
