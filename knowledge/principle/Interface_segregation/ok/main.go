package main

import (
	"fmt"
)

func main() {
	s := impl{}
	s.Buy()
	s.Mail()

	fmt.Println()

	var bs BuyService
	bs = s
	bs.Buy()

	var ms MailService
	ms = s
	ms.Mail()
}

type BuyService interface {
	Buy()
}

type MailService interface {
	Mail()
}

type impl struct{}

func (i impl) Buy() { fmt.Println("Buy") }

func (i impl) Mail() { fmt.Println("Mail") }
