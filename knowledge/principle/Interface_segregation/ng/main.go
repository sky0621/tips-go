package main

import (
	"fmt"
)

func main() {
	var s BuyAndMailService = impl{}
	s.Buy()
	s.Mail()
}

type BuyAndMailService interface {
	Buy()
	Mail()
}

type impl struct{}

func (i impl) Buy() { fmt.Println("Buy") }

func (i impl) Mail() { fmt.Println("Mail") }
