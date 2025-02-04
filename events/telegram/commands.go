package telegram

import (
	"log"
	"os"
	"strings"

	"github.com/EvgeniyGnetnev/Tg-bot_DeepSeek/clients/deepseek"
	"github.com/EvgeniyGnetnev/Tg-bot_DeepSeek/lib/e"
	"github.com/joho/godotenv"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isQuestion(text) {
		return p.sendQuestion(chatID, text, username)
	}

	switch text {
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}


func (p *Processor) sendQuestion(chatID int, text string, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do comand: send question", err) }()

	if err := godotenv.Load(); err != nil {
        log.Fatal("No .env file found")
    }

	ds := deepseek.New(
		os.Getenv("DEEPSEEK_API_KEY"),
		"https://openrouter.ai/api/v1",
	)

	if err := p.tg.SendMessage(chatID, "Думаю..."); err != nil {
		return err
	}

	answer, err := ds.DoRequest(text)
	if err != nil{
		return err
	}

	if err := p.tg.SendMessage(chatID, answer); err != nil {
		return err
	}

	return nil
}


func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func isQuestion(text string) bool {
	return len(text) > 0 && text[0] != '/'
}