package main

import "log"

type talker struct {
	name string
	talk chan string
	recv chan string
	room *talkroom
}

func newTalker(name string, room *talkroom) *talker {
	return &talker{
		name: name,
		talk: make(chan string),
		recv: make(chan string),
		room: room,
	}
}

func (t *talker) talking() {
	log.Printf("<talker.go>[talking][%s] トークルームに入室します。\n", t.name)
	t.room.joinning <- t // トークルームに入室
	log.Printf("<talker.go>[talking][%s] トークルームに入室しました。\n", t.name)
	log.Println("　")

	// トーク中の状態をスタート
	for {
		select {
		case talk := <-t.talk:
			if talk == "" {
				continue
			}
			log.Printf("<talker.go>[talking][%s] トーク「%s」の通知を受けました。\n", t.name, talk)
			// 話者からのトークがあったら、入室しているトークルームに通知
			log.Printf("<talker.go>[talking][%s] トークルームにトークがあった旨を通知します。\n", t.name)
			t.room.talk <- talk
			log.Printf("<talker.go>[talking][%s] トークルームにトークがあった旨を通知しました。\n", t.name)
			log.Println("　")
		case talk := <-t.recv:
			if talk == "" {
				continue
			}
			log.Printf("<talker.go>[talking][%s] 誰かがトーク「%s」をしたため、トークルームから通知が来ました。\n", t.name, talk)
			log.Println("　")
		}
	}
}

func (t *talker) bye() {
	log.Printf("<talker.go>[bye][%s] トークルームから退室します。\n", t.name)
	t.room.leaving <- t // トークルームから退室
	log.Printf("<talker.go>[bye][%s] トークルームから退室しました。\n", t.name)
	close(t.talk)
	close(t.recv)
	log.Println("　")
}
