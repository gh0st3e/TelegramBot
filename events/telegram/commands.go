package telegram

import (
	"fmt"
	"log"
	"os"
)

const (
	HelpCmd                   = "/help"
	StartCmd                  = "/start"
	PreviewOrder              = "/previeworder"
	LogoOrder                 = "/logoorder"
	MontageOrder              = "/montageorder"
	FullChannelManagmentOrder = "/fullchannelmanagmentorder"
)

var ServicesList = map[int]string{
	1: "Логотип",
	2: "Превью",
	3: "Монтаж",
	4: "Полный Менеджмент Канала",
}

var OrderTypeID int = 0

func (p *Processor) doCmd(text string, chatID int, username string) error {

	log.Printf("get new command '%s' from '%s'", text, username)

	switch text {
	case LogoOrder:
		return p.LogoOrder(chatID, username)
	case PreviewOrder:
		return p.PreviewOrder(chatID, username)
	case MontageOrder:
		return p.MontageOrder(chatID, username)
	case FullChannelManagmentOrder:
		return p.FCH_Order(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)

	}
}

func (p *Processor) LogoOrder(chatID int, username string) error {
	OrderTypeID = 1
	SendMail(OrderTypeID, username)
	return p.tg.SendMessage(chatID, msgSendedOrder)
}

func (p *Processor) PreviewOrder(chatID int, username string) error {
	OrderTypeID = 2
	SendMail(OrderTypeID, username)
	return p.tg.SendMessage(chatID, msgSendedOrder)
}

func (p *Processor) MontageOrder(chatID int, username string) error {
	OrderTypeID = 3
	SendMail(OrderTypeID, username)
	return p.tg.SendMessage(chatID, msgSendedOrder)
}

func (p *Processor) FCH_Order(chatID int, username string) error {
	OrderTypeID = 4
	SendMail(OrderTypeID, username)
	return p.tg.SendMessage(chatID, msgSendedOrder)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func SendMail(OrderTypeID int, username string) {

	data := fmt.Sprintf("Пользователь: %s только что сделал заказ : %s \n", username, ServicesList[OrderTypeID])
	file, err := os.OpenFile("events/telegram/orders.txt", os.O_APPEND, 0666)
	log.Printf(data)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(data)

}
