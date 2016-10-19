package main

import (
	"log"
	"sync"
)

// appStopChannel ... アプリ終了を（誰かに対して）通知するためのチャネル
var appStopChannel = make(chan struct{}, 1)

var stopLock sync.Mutex
var isStop = false

// closeApp ...
func closeApp() {
	log.Println("[CloseApp]START")
	log.Println("[CloseApp]アプリ終了フラグ：", isStop)
	stopLock.Lock()
	log.Println("[CloseApp]locked!")
	isStop = true
	stopLock.Unlock()
	log.Println("[CloseApp]unlocked!")
	log.Println("[CloseApp]アプリ終了フラグ：", isStop)
	appStopChannel <- struct{}{}
	log.Println("[CloseApp]END")
}
