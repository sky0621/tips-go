package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{
		"４年,ねんざ,グラウンド",
		"１年,すりきず,ろうか",
		"１年,切りきず,教室",
		"４年,つき指,総合ホール",
		"３年,ねんざ,ろうか",
		"４年,打ぼく,グラウンド",
		"４年,すりきず,総合ホール",
		"１年,すりきず,グラウンド",
		"４年,打ぼく,ろうか",
		"４年,切りきず,教室",
		"３年,打ぼく,グラウンド",
		"３年,すりきず,教室",
		"２年,ねんざ,総合ホール",
		"２年,すりきず,グラウンド",
		"４年,打ぼく,教室",
		"３年,すりきず,ろうか",
		"２年,すりきず,教室",
		"３年,切りきず,グラウンド",
		"４年,切りきず,教室",
		"３年,つき指,グラウンド",
		"３年,打ぼく,総合ホール",
		"１年,すりきず,教室",
		"３年,ねんざ,総合ホール",
		"２年,すりきず,ろうか",
		"３年,切りきず,総合ホール",
		"４年,打ぼく,総合ホール",
		"４年,すりきず,グラウンド",
	}

	gakunen := make(map[string]int)
	syurui := make(map[string]int)
	basho := make(map[string]int)
	for _, i := range input {
		ins := strings.Split(i, ",")
		if v, ok := gakunen[ins[0]]; ok {
			gakunen[ins[0]] = v + 1
		} else {
			gakunen[ins[0]] = 1
		}
		if v, ok := syurui[ins[1]]; ok {
			syurui[ins[1]] = v + 1
		} else {
			syurui[ins[1]] = 1
		}
		if v, ok := basho[ins[2]]; ok {
			basho[ins[2]] = v + 1
		} else {
			basho[ins[2]] = 1
		}
	}

	fmt.Println()

	fmt.Println("【けがをした学年】")
	gakunenTotal := 0
	for _, v := range gakunen {
		gakunenTotal += v
	}
	fmt.Println(gakunen)
	fmt.Printf("合計：%d\n", gakunenTotal)
	fmt.Println()

	fmt.Println("【けがのしゅるい】")
	syuruiTotal := 0
	for _, v := range syurui {
		syuruiTotal += v
	}
	fmt.Println(syurui)
	fmt.Printf("合計：%d\n", syuruiTotal)
	fmt.Println()

	fmt.Println("【けがをした場所】")
	bashoTotal := 0
	for _, v := range basho {
		bashoTotal += v
	}
	fmt.Println(basho)
	fmt.Printf("合計：%d\n", bashoTotal)
	fmt.Println()
}
