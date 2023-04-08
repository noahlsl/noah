package telegramx

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Cfg) NewClient() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		panic(err)
	}
	return bot
}
