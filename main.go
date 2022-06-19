package main

import (
	"flag"
	tgClient "github.com/gh0st3e/OrderBot/clients/telegram"
	event_consumer "github.com/gh0st3e/OrderBot/consumer/event-consumer"
	"github.com/gh0st3e/OrderBot/events/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token os not specified")
	}

	return *token
}
