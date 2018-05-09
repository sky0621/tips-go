package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// SignalChannel ... OSからのシグナル受信用チャネル
var signalChannel = make(chan os.Signal, 1)

type signalReceiver struct {
}

// newSignalReceiver2 ... OSからのSIGINT/SIGTERM着信をチャネルに通知する設定をし、SignalReceiver構造体を初期化してポインタを返す。
func newSignalReceiver2() *signalReceiver {
	log.Println("[NewSignalReceiver]START")
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	log.Println("[NewSignalReceiver]signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)")
	log.Println("[NewSignalReceiver]END")
	return &signalReceiver{}
}

// receiveSignal ... OSからのシグナルを受信したら、引数で渡されたアプリ終了用関数を実行する。
func (sr *signalReceiver) receiveSignal(closeApp func()) {
	log.Println("[ReceiveSignal]START")
	<-signalChannel // OSからのシグナルを待機
	log.Println("[ReceiveSignal]OSからシグナル受信！！！ -> closeAppをコールします。")
	closeApp()
	log.Println("[ReceiveSignal]END")
}
