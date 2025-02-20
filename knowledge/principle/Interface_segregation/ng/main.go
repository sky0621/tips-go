package main

import (
	"fmt"
)

func main() {
	s := impl{}
	s.Buy()
	s.Mail()

	fmt.Println()

	var bms BuyAndMailService
	bms = s
	bms.Buy()
	bms.Mail()
}

type BuyAndMailService interface {
	Buy()
	Mail()
}

type impl struct{}

func (i impl) Buy() { fmt.Println("Buy") }

func (i impl) Mail() { fmt.Println("Mail") }
