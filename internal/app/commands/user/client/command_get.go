package client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *UserClientCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	clientID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong id", args[0])

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "clientID must be unsigned integer")
		c.sendMessage(msg)

		return
	}

	client, err := c.clientService.Describe(clientID)

	if err != nil {
		log.Printf(err.Error())

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.sendMessage(msg)

		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, client.String())
	c.sendMessage(msg)
}
