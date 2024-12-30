package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {

	getwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := filepath.WalkDir(getwd, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
