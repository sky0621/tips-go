package main

import (
	"fmt"
)

func main() {
	var s BuyAndMailService = impl{}
	s.Buy()
	s.Mail()
}

type BuyService interface {
	Buy()
}

type MailService interface {
	Mail()
}

type BuyAndMailService interface {
	BuyService
	MailService
}

type impl struct{}

func (i impl) Buy() { fmt.Println("Buy") }

func (i impl) Mail() { fmt.Println("Mail") }
