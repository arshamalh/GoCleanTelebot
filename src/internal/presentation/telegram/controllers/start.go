package controllers

import (
	"coryptex.com/bot/vip-signal/internal/presentation/telegram/structs"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Start handler for /start command on the bot
func Start(tlg *structs.Telegram, m *tb.Message) {
	tlg.Bot.Send(m.Sender, "Hello "+m.Sender.FirstName+"\nClick /help for help!")
}
