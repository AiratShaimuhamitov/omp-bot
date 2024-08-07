package client

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *UserClientCommander) Edit(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.CommandArguments(), " ")

	if len(args) != 2 {
		log.Println("wrong arguments count")

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "must be 2 arguments separated by space")
		c.sendMessage(msg)
		return
	}

	clientID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		log.Println("wrong id", args[0])

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "clientID must be unsigned integer")
		c.sendMessage(msg)

		return
	}

	inputData := ClientInput{}

	err = json.Unmarshal([]byte(args[1]), &inputData)

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

	err = c.clientService.Update(clientID, c.convertToEntity(inputData))

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
