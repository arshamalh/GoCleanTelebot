package helpers

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/presentation/telegram/structs"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

// CheckAdmin middleware to check if request is from an admin?
func CheckAdmin(ctx context.Context, t *structs.Telegram, m *tb.Message) bool {
	isAdmin, err := t.Authsvc.Authorize(ctx, m.Sender.Recipient())
	if err != nil {
		fmt.Println("Error happened checking admin", err)
		return false
	}
	return isAdmin
}
