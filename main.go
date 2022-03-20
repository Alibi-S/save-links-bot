package main

import (
	"flag"
	"log"

	"github.com/Alibi-S/save-links-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
}

func mustToken() string {
	token := flag.String("telegram-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
