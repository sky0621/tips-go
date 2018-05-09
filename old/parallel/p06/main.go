package main

import (
	"fmt"
	"time"
)

func main() {
	// アカウントを４つ用意
	accs := []*Account{
		&Account{ID: 1, Name: "アカウント１"},
		&Account{ID: 2, Name: "アカウント２"},
		&Account{ID: 3, Name: "アカウント３"},
	}

	// アカウント毎のキューを用意
	queues := []*AccountTaskQueue{}
	for _, acc := range accs {
		queues = append(queues, NewTaskQueue(acc))
	}

	// キューの一つ一つにID:1～9までのタスクをブロードキャスト　⇒　自分が担当するIDの分だけ処理
	for _, q := range queues {
		fmt.Println("queue loop")

		go func() {
			time.Sleep(3 * time.Second)
			tsk := q.Pop()
			tsk.Do()
		}()

		id := 0
		for id < 5 {
			id = id + 1
			fmt.Printf("push %d\n", id)
			go q.Push(&Task{ID: id})
		}

	}

}

type Account struct {
	ID   int
	Name string
}

type Task struct {
	ID int
}

func (t *Task) Do() {
	if t == nil {
		fmt.Println("Task is nil")
		return
	}
	fmt.Printf("%s is doing.", t.ID)
}

type AccountTaskQueue struct {
	acc *Account
	ch  chan *Task
}

func NewTaskQueue(acc *Account) *AccountTaskQueue {
	return &AccountTaskQueue{acc: acc, ch: make(chan *Task)}
}

func (q *AccountTaskQueue) Push(t *Task) {
	for {
		select {
		case q.ch <- t:
			fmt.Println("[Push] case q,ch <- t")
		default:
		}
	}
}

func (q *AccountTaskQueue) Pop() *Task {
	for {
		select {
		case t := <-q.ch:
			fmt.Println("[Pop] case t := <-q.ch")
			return t
		default:
			return nil
		}
	}
}
