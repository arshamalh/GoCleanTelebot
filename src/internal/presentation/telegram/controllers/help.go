package controllers

import (
	"coryptex.com/bot/vip-signal/internal/presentation/telegram/structs"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Help handler for /help command on the bot
func Help(tlg *structs.Telegram, m *tb.Message) {
	tlg.Bot.Send(m.Sender, "Send your .xlsx or .csv file to me")
}
