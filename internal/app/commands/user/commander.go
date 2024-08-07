package user

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/user/client"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type UserCommander struct {
	bot             *tgbotapi.BotAPI
	clientCommander Commander
}

func NewUserCommander(bot *tgbotapi.BotAPI) *UserCommander {
	return &UserCommander{
		bot:             bot,
		clientCommander: client.NewUserClientCommander(bot),
	}
}

func (c *UserCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "client":
		c.clientCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("UserCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *UserCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "client":
		c.clientCommander.HandleCommand(message, commandPath)
	default:
		log.Printf("UserCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
