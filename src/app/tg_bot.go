package app

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Work() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Panic(err, "trouble token tg_bot")
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}

		text := update.Message.Text

		switch text {
		case "/start":
			{
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!\nВведи время на которое завести будильник ")
				bot.Send(msg)
			}
		case "/newtime":
			{
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!\nВведи время на которое завести будильник ")
				bot.Send(msg)
			}

		}
	}
}
