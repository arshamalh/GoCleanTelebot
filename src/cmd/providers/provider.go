package providers

import (
	"context"
	"coryptex.com/bot/vip-signal/pkg/validation"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/tucnak/telebot.v2"
)

// ProvideMongoDB : Provide an MongoDB instance
func ProvideMongoDB() *mongo.Database {
	var constr string
	if os.Getenv("DB_ACCESS_CONTROL") == "" {
		dbUser := os.Getenv("DB_USERNAME")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		validation.PanicIfEmpty(dbUser, "Error loading DB username")
		validation.PanicIfEmpty(dbPass, "Error loading DB password")
		validation.PanicIfEmpty(dbHost, "Error loading DB Host")
		validation.PanicIfEmpty(dbPort, "Error loading DB Port")
		constr = "mongodb://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort
	} else {
		constr = "mongodb://localhost:27017"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(constr))
	if err != nil {
		fmt.Println("Error making new MongoDB client", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("Error connecting to MongoDB client", err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		fmt.Println("Error Pinging to MongoDB client", err)
	}
	return client.Database(os.Getenv("DB_NAME"))
}

// ProvideTelegramHTTPClient : provide a http client to connect with telegram.
func ProvideTelegramHTTPClient() *http.Client {
	return &http.Client{}
}

// ProvideTeleBot : provide an instance of telebot
func ProvideTeleBot() *telebot.Bot {
	botToken := os.Getenv("TOKEN")
	validation.PanicIfEmpty(botToken, "Error loading bot token")
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		fmt.Println("Error happened when initializing the bot with token...", err)
	}
	return bot
}
