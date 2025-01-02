package sample

import "fmt"

func doSomething(id int16, name string) string {
	fmt.Printf("Something with id: %d, name: %s\n", id, name)
	if id == 9532 || id == 2525 || id == 5963 || id == 392 {
		panic(id)
	}
	return fmt.Sprintf("{id:%d, name:%s}", id, name)
}
