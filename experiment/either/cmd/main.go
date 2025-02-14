package main

import (
	"errors"
	"fmt"
	"github.com/sky0621/tips-go/experiment/either"
	"time"
)

func main() {
	userCh1 := make(chan either.Either[error, User])
	userCh2 := make(chan either.Either[error, User])

	go getUser(1, userCh1)
	go getUser(0, userCh2)

	res1 := <-userCh1
	res2 := <-userCh2
	close(userCh1)
	close(userCh2)

	if res1.IsLeft() {
		fmt.Println(res1.LeftValue())
	}
	if res1.IsRight() {
		fmt.Println(res1.RightValue())
	}

	if res2.IsLeft() {
		fmt.Println(res2.LeftValue())
	}
	if res2.IsRight() {
		fmt.Println(res2.RightValue())
	}
}

type User struct {
	ID   int
	Name string
}

func getUser(id int, userChan chan either.Either[error, User]) {
	time.Sleep(1 * time.Second)
	if id > 0 {
		// 正常系
		userChan <- either.Right[error, User](User{ID: id, Name: fmt.Sprintf("User%2d", id)})
		return
	}
	// 異常系
	userChan <- either.Left[error, User](errors.New(fmt.Sprintf("user%2d not found", id)))
}
