package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("MOV_0578.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileInfo, serr := file.Stat()
	if serr != nil {
		fmt.Println("Stat Err!")
		fmt.Println(serr)
		return
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.ModTime())

	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
	fmt.Printf("%+v\n", fileInfo)
	fmt.Println("--------------------------------------------")
	fmt.Printf("%#v\n", fileInfo)
}
