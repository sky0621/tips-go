package sub03

import "fmt"

type App struct {
	S string
}

func (a *App) Print() {
	fmt.Println(a.S)
}
