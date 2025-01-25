package config

import (
	"os"
	"strconv"
)

type Config struct {
	TelegramToken string
	ChannelID     int64
}

func Load() (*Config, error) {
	ID, _ := strconv.Atoi(os.Getenv("TELEGRAM_CHANNEL_ID"))
	return &Config{
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		ChannelID:     int64(ID),
	}, nil
}
