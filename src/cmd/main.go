package main

import (
	"context"
	"coryptex.com/bot/vip-signal/cmd/providers"
	"coryptex.com/bot/vip-signal/internal/data/repositories"
	"coryptex.com/bot/vip-signal/internal/domain/authing"
	"os"
	"os/signal"

	"coryptex.com/bot/vip-signal/internal/data/datasource/mongo"
	"coryptex.com/bot/vip-signal/internal/data/datasource/telegram"
	"coryptex.com/bot/vip-signal/internal/domain/publishing"
	pretel "coryptex.com/bot/vip-signal/internal/presentation/telegram"
)

func main() {
	mdb := providers.ProvideMongoDB()
	teleclient := providers.ProvideTelegramHTTPClient()
	telebot := providers.ProvideTeleBot()

	//** Data sources
	mongodb := mongo.NewMongo(mdb)
	telepub := telegram.NewTelegram(teleclient)

	//** Repositories and Publishers
	spub := repositories.NewSignalPublisher(telepub)
	srepo := repositories.NewSignalRepository(mongodb)
	arepo := repositories.NewAdminRepository(mongodb)

	//** Services
	psvc := publishing.NewService(srepo, spub)
	asvc := authing.NewService(arepo)

	//** Servers
	telesrv := pretel.NewTelegram(telebot, psvc, asvc)

	// ...
	telesrv.Start(context.Background())

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch

	telesrv.Stop(context.Background())
}
