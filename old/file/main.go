package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	const dirn = "C:\\temp\\gofilesrename"
	fileInfos, err := ioutil.ReadDir(dirn)
	if err != nil {
		log.Println(err)
		return
	}
	for _, fileInfo := range fileInfos {
		renamed := strings.Replace(fileInfo.Name(), "_", "", -1)
		nerr := os.Rename(filepath.Join(dirn, fileInfo.Name()), filepath.Join(dirn, renamed))
		if nerr != nil {
			log.Println(nerr)
			return
		}
	}
	log.Println("fin")
}
