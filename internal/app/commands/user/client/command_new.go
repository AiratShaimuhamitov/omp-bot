package client

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *UserClientCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	inputData := ClientInput{}

	err := json.Unmarshal([]byte(args), &inputData)

	if err != nil {
		log.Println("wrong args", args)

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "invalid client data")
		c.sendMessage(msg)

		return
	}

	err = c.validate(inputData)
	if err != nil {
		log.Println("validation error", err.Error())

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.sendMessage(msg)

		return
	}

	clientID, err := c.clientService.Create(c.convertToEntity(inputData))

	if err != nil {
		log.Println("create error", err.Error())

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "error create client, try again")
		c.sendMessage(msg)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("client created with ID %d", clientID),
	)

	c.sendMessage(msg)
}
