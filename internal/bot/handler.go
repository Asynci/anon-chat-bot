package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func handlerMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update, channel int64) {
	if update.Message.Photo != nil {
		handlerPhoto(bot, update.Message, channel)
	}
	if update.Message.Video != nil {
		handlerVideo(bot, update.Message, channel)
	}
	if update.Message.Document != nil {
		handlerDocument(bot, update.Message, channel)
	}
	if update.Message.Audio != nil {
		handlerAudio(bot, update.Message, channel)
	}
	if update.Message.Text != "" {
		handlerText(bot, update.Message, channel)
	}
}

func handlerPhoto(bot *tgbotapi.BotAPI, message *tgbotapi.Message, channel int64) {
	photo := message.Photo[len(message.Photo)-1]

	msg := tgbotapi.NewPhoto(channel, tgbotapi.FileID(photo.FileID))
	caption := message.Caption
	if caption == "" {
		caption = "Новое фото 📸"
	}

	msg.Caption = fmt.Sprintf(
		"<b>%s</b>\n\n<a href=\"https://t.me/anon_malmyzh_bot\">📭 Написать анонимное сообщение</a>",
		caption,
	)
	msg.ParseMode = "HTML"

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки фото: %v", err)
	}

	msgTo := tgbotapi.NewMessage(message.Chat.ID, "<b>Сообщение отправлено 💌</b>")
	msgTo.ParseMode = "HTML"
	msgTo.ReplyToMessageID = message.MessageID
	_, err = bot.Send(msgTo)
	if err != nil {
		log.Printf("Ошибка")
	}
}

func handlerVideo(bot *tgbotapi.BotAPI, message *tgbotapi.Message, channel int64) {
	video := message.Video

	msg := tgbotapi.NewVideo(channel, tgbotapi.FileID(video.FileID))
	caption := message.Caption
	if caption == "" {
		caption = "Новое видео 🎥"
	}

	msg.Caption = fmt.Sprintf(
		"<b>%s</b>\n\n<a href=\"https://t.me/anon_malmyzh_bot\">📭 Написать анонимное сообщение</a>",
		caption,
	)
	msg.ParseMode = "HTML"

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки видео: %v", err)
	}

	msgTo := tgbotapi.NewMessage(message.Chat.ID, "<b>Сообщение отправлено 💌</b>")
	msgTo.ParseMode = "HTML"
	msgTo.ReplyToMessageID = message.MessageID
	_, err = bot.Send(msgTo)
	if err != nil {
		log.Printf("Ошибка")
	}
}

func handlerAudio(bot *tgbotapi.BotAPI, message *tgbotapi.Message, channel int64) {
	audio := message.Audio
	msg := tgbotapi.NewAudio(channel, tgbotapi.FileID(audio.FileID))

	caption := msg.Caption
	msg.Caption = fmt.Sprintf(
		"<b>%s</b>\n\n<a href=\"https://t.me/anon_malmyzh_bot\">📭 Написать анонимное сообщение</a>",
		caption,
	)
	msg.ParseMode = "HTML"

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки аудио: %v", err)
	}

	msgTo := tgbotapi.NewMessage(message.Chat.ID, "<b>Сообщение отправлено 💌</b>")
	msgTo.ParseMode = "HTML"
	msgTo.ReplyToMessageID = message.MessageID
	_, err = bot.Send(msgTo)
	if err != nil {
		log.Printf("Ошибка")
	}
}

func handlerText(bot *tgbotapi.BotAPI, message *tgbotapi.Message, channel int64) {
	text := fmt.Sprintf(
		"<b>Новое сообщение 📩:</b>\n\n%s\n\n<a href=\"https://t.me/anon_malmyzh_bot\">📭 Написать анонимное сообщение</a>",
		message.Text,
	)

	msg := tgbotapi.NewMessage(channel, text)
	msg.ParseMode = "HTML"

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("error sending message: %v", err)
	}

	msgTo := tgbotapi.NewMessage(message.Chat.ID, "<b>Сообщение отправлено 💌</b>")
	msgTo.ParseMode = "HTML"
	msgTo.ReplyToMessageID = message.MessageID
	_, err = bot.Send(msgTo)
	if err != nil {
		log.Printf("Ошибка")
	}
}

func handlerDocument(bot *tgbotapi.BotAPI, message *tgbotapi.Message, channel int64) {
	doc := message.Document
	msg := tgbotapi.NewDocument(channel, tgbotapi.FileID(doc.FileID))

	caption := message.Caption
	if caption == "" {
		caption = "Новый документ 📄"
	}

	msg.Caption = fmt.Sprintf(
		"<b>%s</b>\n\n<a href=\"https://t.me/anon_malmyzh_bot\">📭 Написать анонимное сообщение</a>",
		caption,
	)
	msg.ParseMode = "HTML"

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки документа: %v", err)
	}

	msgTo := tgbotapi.NewMessage(message.Chat.ID, "<b>Сообщение отправлено 💌</b>")
	msgTo.ParseMode = "HTML"
	msgTo.ReplyToMessageID = message.MessageID
	_, err = bot.Send(msgTo)
	if err != nil {
		log.Printf("Ошибка")
	}
}
