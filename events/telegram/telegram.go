package telegram

import "github.com/EvgeniyGnetnev/Tg-bot_DeepSeek/clients/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client)