package client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const helpText = `
/help__user__client — print list of commands
/get__user__client <ID> — get a entity
/list__user__client — get a list of your entity
/delete__user__client <ID> — delete an existing entity
/new__user__client <JSON> — create a new entity
/edit__user__client <JSON> — edit a entity

    <ID> - unsigned integer
    <JSON> - client data serialized in JSON format
    Example <JSON>: {"firstName":"Kevin","secondName":"Trump"}
`

func (c *UserClientCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, helpText)

	c.sendMessage(msg)
}
