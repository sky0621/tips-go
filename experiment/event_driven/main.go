package main

import (
	"fmt"
	"time"
)

type Kind int

const (
	Add Kind = iota
	Edit
	Delete
)

type Event struct {
	Kind      Kind
	Data      any
	Timestamp string
}

type EventBus struct {
	subscribers map[Kind][]chan<- Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[Kind][]chan<- Event),
	}
}

func (b *EventBus) Subscribe(kind Kind, subscriber chan<- Event) {
	b.subscribers[kind] = append(b.subscribers[kind], subscriber)
}

func (b *EventBus) Publish(event Event) {
	targets := b.subscribers[event.Kind]
	for _, subscriber := range targets {
		subscriber <- event
	}
}

func show(ev chan Event) {
	for e := range ev {
		fmt.Println(e)
	}
}

func main() {
	eventBus := NewEventBus()

	userAdded := make(chan Event)
	userEdited := make(chan Event)
	userDeleted := make(chan Event)

	go show(userAdded)
	go show(userEdited)
	go show(userDeleted)

	eventBus.Subscribe(Add, userAdded)
	eventBus.Subscribe(Edit, userEdited)
	eventBus.Subscribe(Delete, userDeleted)

	eventBus.Publish(Event{
		Kind: Add, Data: "user A added", Timestamp: time.Now().Format(time.RFC3339Nano),
	})
	time.Sleep(100 * time.Millisecond)
	eventBus.Publish(Event{
		Kind: Edit, Data: "user A edited", Timestamp: time.Now().Format(time.RFC3339Nano),
	})
	time.Sleep(100 * time.Millisecond)
	eventBus.Publish(Event{
		Kind: Delete, Data: "user A deleted", Timestamp: time.Now().Format(time.RFC3339Nano),
	})
	time.Sleep(300 * time.Millisecond)

	eventBus.Publish(Event{
		Kind: Add, Data: "user B added", Timestamp: time.Now().Format(time.RFC3339Nano),
	})
	time.Sleep(100 * time.Millisecond)
	eventBus.Publish(Event{
		Kind: Edit, Data: "user B edited", Timestamp: time.Now().Format(time.RFC3339Nano),
	})
	time.Sleep(250 * time.Millisecond)

	eventBus.Publish(Event{
		Kind: Add, Data: "user C added", Timestamp: time.Now().Format(time.RFC3339Nano),
	})
	time.Sleep(100 * time.Millisecond)
	eventBus.Publish(Event{
		Kind: Delete, Data: "user C deleted", Timestamp: time.Now().Format(time.RFC3339Nano),
	})

	time.Sleep(3 * time.Second)
}
