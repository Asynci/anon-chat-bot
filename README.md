# Anon Chat Bot

Telegram-бот для анонимной переписки. Пользователи могут отправлять сообщения в указанный канал, сохраняя анонимность. Бот поддерживает текстовые сообщения, фото, видео, аудио и документы.

---

## Что нужно для запуска

1. **Telegram Bot API токен** — создайте бота через [BotFather](https://core.telegram.org/bots#botfather) и получите токен.
2. **ID канала** — создайте канал в Telegram и получите его ID.
3. **Go** — установите Go версии 1.22 или выше.

---

## Как запустить

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/Asynci/anon-chat-bot.git
   cd anon-chat-bot
2. Установите зависимости:
   ```bash
   go mod tidy
3. Создайте переменные окружения:
   ```bash
   export TELEGRAM_BOT_TOKEN="your_token"
   export TELEGRAM_CHANNEL_ID="ypur_id (example=12344554)"
4. Соберите проект:
   ```bash
   go build -o bin/bot ./cmd/app/main.go
5. Запустите бота:
   ```bash
   ./bin/bot