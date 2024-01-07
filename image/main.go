package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(-1)
	}

	fromDir := os.Args[1]
	toDir := os.Args[2]

	fileList, err := os.Create(filepath.Join(toDir, "fileList.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := fileList.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	existsSet := mapset.NewSet[string]()

	fileListScanner := bufio.NewScanner(fileList)
	for fileListScanner.Scan() {
		existsSet.Add(fileListScanner.Text())
	}

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

		return exec(path, existsSet, toDir, fi)
	}); err != nil {
		log.Fatal(err)
	}

	for _, s := range existsSet.ToSlice() {
		_, err := fileList.WriteString(fmt.Sprintf("%s\n", s))
		if err != nil {
			log.Println(err)
		}
	}
}

func exec(path string, existsSet mapset.Set[string], toDir string, fi fs.FileInfo) error {
	if !addExists(existsSet, fi) {
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

	outFile, err := os.Create(filepath.Join(toDir, createOutFileName(fi)))
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
}

func addExists(existsSet mapset.Set[string], fi fs.FileInfo) bool {
	key := fi.Name() + formatModTime(fi.ModTime())
	if existsSet.Contains(key) {
		log.Println("=====================")
		log.Println("EXISTS:", key)
		log.Println("=====================")
		return false
	}
	existsSet.Add(key)
	return true
}

func formatModTime(modTime time.Time) string {
	return modTime.Format("2006-01-02T15h04m05s")
}

func createOutFileName(fi fs.FileInfo) string {
	return fmt.Sprintf("%s%s", formatModTime(fi.ModTime()), filepath.Ext(fi.Name()))
}
