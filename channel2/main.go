package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// OSシグナル受信時にコールされるまでクローズ関数を待機させるためのチャネル
	stopChan := make(chan int, 1)

	// OSシグナル待ち受けゴルーチン起動
	r := NewSignalReceiverImpl(stopChan)
	go r.ReceiveSignal(func(stopChan chan int) {
		fmt.Println("Finish!!!")
	})(stopChan)
}

type CloseFunc func(stopChan chan int)

type SignalReceiver interface {
	ReceiveSignal(fn CloseFunc)
}

type SignalReceiverImpl struct {
	signalChan chan os.Signal
	stopChan chan int
}

func NewSignalReceiverImpl(stopChan chan int) SignalReceiver {
	r := &SignalReceiverImpl{
		signalChan: make(chan os.Signal, 1),
		stopChan: stopChan,
	}
	signal.Notify(r.signalChan, syscall.SIGINT, syscall.SIGTERM)
	return r
}

func (r *SignalReceiverImpl) ReceiveSignal(fn CloseFunc) {
	<-r.signalChan
	fn()
}
