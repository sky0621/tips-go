package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// OSシグナル待ち受けゴルーチン起動
	r := NewSignalReceiverImpl()
	go func() {
		r.ReceiveSignal(func(stopChan chan int) {
			// アプリ停止時用のチャネルへの受信を待機
			<-stopChan
			fmt.Println("Stop!!!")
		})
	}()
}

type CloseFunc func(stopChan chan int)

type SignalReceiver interface {
	ReceiveSignal(fn CloseFunc)
}

type SignalReceiverImpl struct {
	signalChan chan os.Signal
	stopChan   chan int
}

func NewSignalReceiverImpl() SignalReceiver {
	r := &SignalReceiverImpl{
		signalChan: make(chan os.Signal, 1),
		stopChan:   make(chan int, 1),
	}
	signal.Notify(r.signalChan, syscall.SIGINT, syscall.SIGTERM)
	return r
}

// OSからのシグナルを受信したら、パラメータであるCloseFunc関数を実行する！
func (r *SignalReceiverImpl) ReceiveSignal(fn CloseFunc) {
	<-r.signalChan
	fn(r.stopChan)
}
