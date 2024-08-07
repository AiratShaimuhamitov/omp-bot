package client

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *UserClientCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	clientID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong id", args[0])

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "clientID must be unsigned integer")
		c.sendMessage(msg)

		return
	}

	_, err = c.clientService.Remove(clientID)

	if err != nil {
		log.Println(err.Error())

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.sendMessage(msg)

		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("client with ID %d successfully deleted", clientID))
	c.sendMessage(msg)
}
