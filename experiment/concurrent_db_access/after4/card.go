package main

type Card struct {
	ID   int
	Name string
}

type ListCardFunc func() []*Card

func ListCard() ListCardFunc {}
