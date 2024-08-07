package client

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *UserClientCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "All clients: \n\n"

	clients, err := c.clientService.List(0, DefaultListLimit)

	if err != nil {
		log.Printf(err.Error())

		return
	}

	for _, p := range clients {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if uint64(len(clients)) == DefaultListLimit {
		serializedData, err := json.Marshal(
			CallbackListData{
				Cursor: DefaultListLimit,
				Limit:  DefaultListLimit,
			},
		)

		if err != nil {
			log.Printf(err.Error())

			return
		}

		callbackPath := path.CallbackPath{
			Domain:       "user",
			Subdomain:    "client",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(NextButtonText, callbackPath.String()),
			),
		)
	}

	c.sendMessage(msg)
}
