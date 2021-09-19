package structs

import (
	"coryptex.com/bot/vip-signal/internal/domain/authing"
	"coryptex.com/bot/vip-signal/internal/domain/publishing"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Telegram services
type Telegram struct {
	Bot     *tb.Bot
	Pubsvc  publishing.Service
	Authsvc authing.Service
}
