package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type signalReceiver struct {
	ch           chan os.Signal // OSからのシグナル受信用チャネル
	notifyToStop func()         // OSからのシグナル受信時の他チャネルへの停止通知関数
}

func newSignalReceiver(notifyToStop func()) *signalReceiver {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	return &signalReceiver{
		ch:           ch,
		notifyToStop: notifyToStop,
	}
}

func (r *signalReceiver) wait() {
	<-r.ch
	fmt.Println("received signal !!!!!")
	r.notifyToStop()
}
