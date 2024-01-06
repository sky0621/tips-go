package main

import (
	"fmt"
	_ "image/gif"
	_ "image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(-1)
	}

	fromDir := os.Args[1]
	log.Println(fromDir)
	toDir := os.Args[2]
	log.Println(toDir)

	existsSet := mapset.NewSet[string]()

	if err := filepath.WalkDir(fromDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(path)

		fi, err := d.Info()
		if err != nil {
			log.Println(err)
			return nil
		}

		if fi.IsDir() {
			return nil
		}

		if strings.HasPrefix(fi.Name(), ".") {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return nil
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Println(err)
			}
		}()

		mTime := fi.ModTime().Format("2006-01-02T15h04m05s")

		key := fi.Name() + mTime
		if existsSet.Contains(key) {
			log.Println("=====================")
			log.Println("EXISTS")
			log.Println(key)
			log.Println("=====================")
			return nil
		}
		existsSet.Add(key)

		ext := filepath.Ext(fi.Name())

		outFileName := fmt.Sprintf("%s%s", mTime, ext)
		log.Println(outFileName)

		outFile, err := os.Create(filepath.Join(toDir, outFileName))
		if err != nil {
			log.Println(err)
			return nil
		}
		defer func() {
			if err := outFile.Close(); err != nil {
				log.Println(err)
			}
		}()

		if _, err := io.Copy(outFile, f); err != nil {
			log.Println(err)
			return nil
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
