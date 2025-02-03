package main

import (
	"log"
	"os"

	tgClient "github.com/EvgeniyGnetnev/Tg-bot_DeepSeek/clients/telegram"
	event_consumer "github.com/EvgeniyGnetnev/Tg-bot_DeepSeek/consumer/event-consumer"
	"github.com/EvgeniyGnetnev/Tg-bot_DeepSeek/events/telegram"
	"github.com/EvgeniyGnetnev/Tg-bot_DeepSeek/storage/files"
	"github.com/joho/godotenv"
)

// "os"

// "github.com/EvgeniyGnetnev/Tg-bot_DeepSeek/clients/telegram"

const (
	storagePath = "storage"
	batchSise   = 100
)

func main() {

	if err := godotenv.Load(); err != nil {
        log.Fatal("No .env file found")
    }

	eventsProcessor := telegram.New(
		tgClient.New("api.telegram.org", os.Getenv("TG_BOT_TOKEN")),
		files.New(storagePath))

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSise)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
