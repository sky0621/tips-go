package main

import "log"

func main() {
	file, err := SetupLog(".", "app.log_with_caller_line")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	log.Println("コンソール、および、ファイルにも出力！")
}
