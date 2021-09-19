package controllers

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"coryptex.com/bot/vip-signal/internal/presentation/telegram/helpers"
	"coryptex.com/bot/vip-signal/internal/presentation/telegram/structs"
	"coryptex.com/bot/vip-signal/pkg/excel"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

// UpFile handle request for uploading file, and use helpers to downloading file.
func UpFile(ctx context.Context, tlg *structs.Telegram, m *tb.Message) {
	if helpers.CheckAdmin(ctx, tlg, m) {
		if helpers.CheckFileType(m) {
			lastMessage, err := tlg.Bot.Send(m.Sender, "Uploading...")
			if err != nil {
				fmt.Println("Error responding user", err)
			}
			filePath, err := helpers.UpFile(m)
			if err != nil {
				fmt.Println("Error  Downloading the file", err)
			}
			ss, err := excel.ReadSignals(filePath, "Signal VIP")
			if err != nil {
				// return nil, err
				fmt.Println("Error reading signals", err)
			}
			var signalsList []entities.Signal
			for _, s := range ss {
				if valid, eMessage := s.IsValid(); valid {
					signalsList = append(signalsList, s)
				} else {
					_, _ = tlg.Bot.Send(m.Sender, "Signal is not valid\n"+"Signal: "+s.ToString()+"\nError: "+eMessage)
				}
			}
			_, err = tlg.Bot.Edit(lastMessage, "Uploaded.\nPublishing signals...")
			if err != nil {
				fmt.Println("Error responding user", err)
			}

			if e := tlg.Pubsvc.Publish(ctx, signalsList); e != nil {
				fmt.Println("Error publishing messages", err)
				if _, err := tlg.Bot.Send(m.Sender, "We Can't publish messages at this moment, please try again later"); err != nil {
					fmt.Println("Error responding to user.")
				}
			} else {
				_, err = tlg.Bot.Edit(lastMessage, "All signals published successfully.")
				if err != nil {
					fmt.Println("Error responding user", err)
				}
			}
		} else {
			_, err := tlg.Bot.Send(m.Sender, "File Type is not allowed.")
			if err != nil {
				fmt.Println("Error responding user", err)
			}
		}
	} else {
		_, err := tlg.Bot.Send(m.Sender, "You are not allowed, you need admin permission.")
		if err != nil {
			fmt.Println("Error responding user", err)
		}
	}
}
