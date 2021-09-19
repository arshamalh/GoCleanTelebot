package controllers

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/presentation/telegram/helpers"
	"coryptex.com/bot/vip-signal/internal/presentation/telegram/structs"
	"fmt"
	"github.com/pkg/errors"
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

// SetAdmin handler for setting in admin on the bot
func SetAdmin(tlg *structs.Telegram, m *tb.Message) {
	if helpers.CheckAdmin(context.TODO(), tlg, m) {
		adminTID := strings.Split(m.Text, " ")[1]
		_, err := tlg.Bot.Send(m.Sender, "Admin "+adminTID+" has set")
		if err != nil {
			fmt.Println("Error Responding User", errors.Wrap(err, "internal.presentation.telegram.controllers.set_admin"))
		}

		err = tlg.Authsvc.NewAdmin(context.Background(), adminTID)
		if err != nil {
			fmt.Println("Error setting Admin", errors.Wrap(err, "internal.presentation.telegram.controllers.set_admin"))
		}
	} else {
		_, err := tlg.Bot.Send(m.Sender, "You are not allowed.")
		if err != nil {
			fmt.Println("Error Responding User", errors.Wrap(err, "internal.presentation.telegram.controllers.set_admin"))
		}
	}
}
