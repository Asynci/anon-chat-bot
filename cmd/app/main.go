package main

import (
	"github.com/Asynci/anon-chat-bot/internal/bot"
	"github.com/Asynci/anon-chat-bot/internal/config"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	bot.Start(cfg)
}
