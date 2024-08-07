package client

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/user"
	"github.com/ozonmp/omp-bot/internal/service/user/client"
	"log"
)

type ClientCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type UserClientCommander struct {
	bot           *tgbotapi.BotAPI
	clientService *client.DummyClientService
}

func NewUserClientCommander(bot *tgbotapi.BotAPI) *UserClientCommander {
	clientService := client.NewDummyClientService()

	return &UserClientCommander{
		bot:           bot,
		clientService: clientService,
	}
}

func (c *UserClientCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ClientCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *UserClientCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(message)
	case "get":
		c.Get(message)
	case "list":
		c.List(message)
	case "delete":
		c.Delete(message)
	case "new":
		c.New(message)
	case "edit":
		c.Edit(message)
	}
}

func (c *UserClientCommander) sendMessage(msg tgbotapi.MessageConfig) {
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("UserClientCommander: error sending reply message to chat - %v", err)
	}
}

func (c *UserClientCommander) validate(input ClientInput) error {
	if len(input.FirstName) < 2 || len(input.FirstName) > 255 {
		return errors.New("firstName must be between 2-255 chars")
	}

	if len(input.SecondName) < 2 || len(input.SecondName) > 255 {
		return errors.New("secondName must be between 2-255 chars")
	}

	return nil
}

func (c *UserClientCommander) convertToEntity(input ClientInput) user.Client {
	return user.Client{
		FirstName:  input.FirstName,
		SecondName: input.FirstName,
	}
}
