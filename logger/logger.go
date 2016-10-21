package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// SetupLog ...
func SetupLog(outputDir string, outputFile string) (*os.File, error) {
	logfile, err := os.OpenFile(filepath.Join(outputDir, outputFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("[%s]のログファイル「%s」オープンに失敗しました。 [ERROR]%s\n", outputDir, outputFile, err)
		return nil, err
	}

	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	return logfile, nil
}
