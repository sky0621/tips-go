package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const basePath = "/Volumes/HD-PCGU3-A/data/videos/___Volumes___HD-LCU3___.Trash-1000___files___20220913___/"
const basePath2 = "/Volumes/HD-PCGU3-A/data/videos/___Volumes___HD-LCU3___AQUOS_SHV40_Backup___20190406___Pictures___YMovie___/"
const basePath3 = "/Users/sky0621/Downloads/"

const file1 = "2022-10-08T16h35m32s_95962c7b-f15a-46e7-a8f3-b3bc8eadaef5_190406_180958.mp4"
const file2 = "2019-04-06T07h35m03s_4168f34c-e694-44d3-b665-046ea06d9954_181219_122902.mp4"

func main() {
	b1 := exec(basePath + file1)
	b2 := exec(basePath3 + file1)
	fmt.Println(bytes.Compare(b1, b2))
	b3 := exec(basePath2 + file2)
	b4 := exec(basePath3 + file2)
	fmt.Println(bytes.Compare(b3, b4))
	fmt.Println(bytes.Compare(b1, b3))
}

type CloseFunc func()

func open(path string) (*os.File, CloseFunc) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return f, func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}
}

func exec(path string) []byte {
	f1, c1 := open(path)
	defer c1()

	byteArray := make([]byte, 1024)
	_, err := f1.Read(byteArray)
	if err != nil {
		log.Fatal(err)
	}

	return byteArray
}
