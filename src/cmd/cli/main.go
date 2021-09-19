package main

import (
	"context"
	"coryptex.com/bot/vip-signal/cmd/providers"
	"coryptex.com/bot/vip-signal/internal/data/datasource/mongo"
	"coryptex.com/bot/vip-signal/internal/data/datasource/telegram"
	"coryptex.com/bot/vip-signal/internal/data/repositories"
	"coryptex.com/bot/vip-signal/internal/domain/authing"
	"coryptex.com/bot/vip-signal/internal/domain/publishing"
	"coryptex.com/bot/vip-signal/internal/presentation/cli"
)

func main() {
	//** third-parties: db, http client, ...
	mdb := providers.ProvideMongoDB()
	teleclient := providers.ProvideTelegramHTTPClient()

	//** Data sources
	mongodb := mongo.NewMongo(mdb)
	telepub := telegram.NewTelegram(teleclient)

	//** Repositories and Publishers
	spub := repositories.NewSignalPublisher(telepub)
	srepo := repositories.NewSignalRepository(mongodb)
	arepo := repositories.NewAdminRepository(mongodb)

	//** Services
	asvc := authing.NewService(arepo)
	psvc := publishing.NewService(srepo, spub)

	//** Start
	cliCmds := cli.NewCLI(psvc, asvc)
	cliCmds.Start(context.Background())
}
