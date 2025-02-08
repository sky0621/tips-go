package main

import (
	"fmt"
)

type PlayFunc func(what string) error

type Player struct {
	Name string
	Play PlayFunc
}

func (p Player) execPlay(what string) error {
	return p.Play(what)
}

func main() {
	p1 := Player{
		Name: "Taro",
		Play: func(what string) error {
			switch what {
			case "カードゲーム":
				fmt.Println("それで遊ぼう！")
				return nil
			case "ボードゲーム":
				fmt.Println("まあ、それでもいいか。")
				return nil
			default:
				return fmt.Errorf("%sでは遊べない。", what)
			}
		},
	}
	if err := p1.execPlay("カードゲーム"); err != nil {
		fmt.Println(err)
	}
	if err := p1.execPlay("ボードゲーム"); err != nil {
		fmt.Println(err)
	}
	if err := p1.execPlay("スポーツ"); err != nil {
		fmt.Println(err)
	}
}
