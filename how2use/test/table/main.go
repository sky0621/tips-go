package main

import "fmt"

func main() {
	var year int
	fmt.Print("4桁の数字（年）を入力してください: ")
	_, err := fmt.Scan(&year)
	if err != nil {
		fmt.Println("入力エラー:", err)
		return
	}
	century, err := getCentury(year)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Printf("%d年は%d世紀です。\n", year, century)
}

func getCentury(year int) (int, error) {
	if year < 1000 || year > 9999 {
		return -1, fmt.Errorf("4桁の数字を入力してください: %d", year)
	}
	return (year + 99) / 100, nil
}
