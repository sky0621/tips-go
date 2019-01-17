package main

import (
	"fmt"
	"strings"
)

type book struct {
	ID     int
	Name   string
	Author string
	Price  int
}

func main() {
	b := book{
		ID:     1,
		Name:   "何かの本",
		Author: "著者A",
		Price:  100,
	}

	builder := NewStringBuilder()
	jsonLog := builder.
		Append(`{`).
		Append(`"book": {`).
		Append(`"id": `).AppendInt(b.ID).Append(", ").
		Append(`"name": `).Append(b.Name).Append(", ").
		Append(`"author": `).Append(b.Author).Append(", ").
		Append(`"price": `).AppendInt(b.Price).Append("").
		Append(`}`).
		Append(`}`).
		ToString()

	fmt.Println(jsonLog)
}

// NewStringBuilder ...
func NewStringBuilder() StringBuilder {
	return &stringBuilder{sb: strings.Builder{}, anyErrors: []error{}}
}

// StringBuilder ...
type StringBuilder interface {
	Append(str string) StringBuilder
	AppendInt(num int) StringBuilder
	ToString() string
	Errors() []error
}

type stringBuilder struct {
	sb        strings.Builder
	anyErrors []error
}

// Append ...
func (b *stringBuilder) Append(str string) StringBuilder {
	_, err := b.sb.WriteString(str)
	if err != nil {
		b.anyErrors = append(b.anyErrors, err)
	}
	return b
}

// AppendInt ...
func (b *stringBuilder) AppendInt(num int) StringBuilder {
	_, err := b.sb.WriteString(fmt.Sprintf("%d", num))
	if err != nil {
		b.anyErrors = append(b.anyErrors, err)
	}
	return b
}

// ToString ...
func (b *stringBuilder) ToString() string {
	if len(b.anyErrors) > 0 {
		return ""
	}
	return b.sb.String()
}

// Errors ...
func (b *stringBuilder) Errors() []error {
	return b.anyErrors
}
