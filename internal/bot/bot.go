package bot

import (
	"github.com/Asynci/anon-chat-bot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Start(cfg *config.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	log.Printf("Start bot: %s", bot.Self.UserName)
	channel := cfg.ChannelID
	log.Print(channel)
	if err != nil {
		log.Print(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		handlerMessage(bot, &update, channel)
	}
}
