package company

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CompanyCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here the companies: \n\n"

	companies, err := c.companyService.List(0, 10)
	if err != nil {
		log.Printf("CompanyCommander.List: Error get compines - %v", err)
	}

	for _, c := range companies {
		outputMsgText += c.String() + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.List: error sending reply message to chat - %v", err)
	}
}
