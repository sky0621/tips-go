package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Println("[main]START")
	fl, err := setupLog()
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		time.Sleep(60 * time.Second)
		fl.Close()
	}()
	signalReceiver := newSignalReceiver()
	go signalReceiver.receiveSignal(closeApp)
	go func() {
		log.Println("[go func()]START")
		// time.
		log.Println("[go func()]END")
	}()
	log.Println("[main]END")
}
