package telegram

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain/authing"
	"coryptex.com/bot/vip-signal/internal/domain/publishing"
	"coryptex.com/bot/vip-signal/internal/presentation/telegram/controllers"
	"coryptex.com/bot/vip-signal/internal/presentation/telegram/structs"

	tb "gopkg.in/tucnak/telebot.v2"
)

// TlgI is Telegram interface
type TlgI interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}

// Telegram services
type Telegram structs.Telegram

// NewTelegram make new instance of telegram to use in services
func NewTelegram(bot *tb.Bot, pubsvc publishing.Service, authsvc authing.Service) TlgI {
	return &Telegram{
		Bot:     bot,
		Pubsvc:  pubsvc,
		Authsvc: authsvc,
	}
}

// Start the telegram bot
func (t *Telegram) Start(ctx context.Context) {
	SetControllers(ctx, (*structs.Telegram)(t))
	t.Bot.Start()
}

// Stop the telegram bot
func (t *Telegram) Stop(ctx context.Context) {
	t.Bot.Stop()
}

// SetControllers will set command handlers on the bot
func SetControllers(ctx context.Context, t *structs.Telegram) {
	t.Bot.Handle("/start", func(m *tb.Message) { controllers.Start(t, m) })
	t.Bot.Handle("/help", func(m *tb.Message) { controllers.Help(t, m) })
	t.Bot.Handle("/setadmin", func(m *tb.Message) { controllers.SetAdmin(t, m) })
	t.Bot.Handle("setadmin", func(m *tb.Message) { controllers.SetAdmin(t, m) })
	t.Bot.Handle(tb.OnDocument, func(m *tb.Message) { controllers.UpFile(ctx, t, m) })
	//t.Bot.Handle("/setcmd", func(m *tb.Message) {})
	/// TODO: t *structs.Telegram have to change to something better.
	/// TODO: IS there any need for Admin to Set commands when there is no handler for that command?
}
