package bot

import (
	"database/sql"
	"github.com/Asynci/anon-chat-bot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

const (
	cooldownDuration = 45 * time.Second
	dbPath           = "bot.db"
)

var db *sql.DB

func Start(cfg *config.Config) {
	log.Printf("Bot start")
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user_cooldowns (
			user_id INTEGER PRIMARY KEY,
			last_message_time DATETIME
		);
	`)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %v", err)
	}

	bot.Debug = true
	log.Printf("Бот запущен: %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if isOnCooldown(update.Message.From.ID) {
			log.Printf("Пользователь %d на кулдауне", update.Message.From.ID)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "<b>⏳ Подождите 45 секунд перед отправкой следующего сообщения.</b>")
			msg.ParseMode = "HTML"
			msg.ReplyToMessageID = update.Message.MessageID
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Ошибка отправки сообщения о кулдауне: %v", err)
			}

			continue
		}

		updateLastMessageTime(update.Message.From.ID)

		handlerMessage(bot, &update, cfg.ChannelID)
	}
}

func isOnCooldown(userID int64) bool {
	var lastMessageTime time.Time
	err := db.QueryRow(`
		SELECT last_message_time
		FROM user_cooldowns
		WHERE user_id = ?
	`, userID).Scan(&lastMessageTime)

	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Printf("Ошибка запроса к базе данных: %v", err)
		return false
	}

	return time.Since(lastMessageTime) < cooldownDuration
}

func updateLastMessageTime(userID int64) {
	_, err := db.Exec(`
		INSERT OR REPLACE INTO user_cooldowns (user_id, last_message_time)
		VALUES (?, ?)
	`, userID, time.Now())
	if err != nil {
		log.Printf("Ошибка обновления времени сообщения: %v", err)
	}
}
