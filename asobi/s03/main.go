package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	const (
		un = "運動場"
		ro = "廊下"
		ky = "教室"
		ta = "体育館"
		ka = "階だん"
	)

	const (
		ki = "切りきず"
		da = "打ぼく"
		su = "すりきず"
		ko = "こっせつ"
		tu = "つきゆび"
		ne = "ねんざ"
	)

	fn := func(a, b string) string {
		return a + "," + b
	}

	input := []string{
		fn(ro, da),
		fn(ta, su),
		fn(ky, su),
		fn(ro, tu),
		fn(ta, su),
		fn(un, tu),
		fn(ta, tu),
		fn(ky, da),
		fn(ta, su),
		fn(un, da),
		fn(ta, ki),
		fn(ky, ki),

		fn(ky, ki),
		fn(un, su),
		fn(ta, ko),
		fn(un, su),
		fn(ro, su),
		fn(ro, ki),
		fn(ka, da),
		fn(un, ne),
		fn(ta, su),
		fn(ky, su),
		fn(ta, tu),
		fn(ky, da),

		fn(ky, ki),
		fn(un, da),
		fn(ka, su),
		fn(un, ki),
		fn(ta, da),
		fn(ky, da),
		fn(ka, da),
		fn(ta, ne),
		fn(ky, su),
		fn(un, su),
		fn(ta, tu),
		fn(ro, da),
	}

	sort.Strings(input)

	bashoMap := make(map[string]int)
	typeMap := make(map[string]int)

	bashoBanpei := ""
	typeBanpei := ""
	for _, i := range input {
		bashoType := strings.Split(i, ",")

		if bashoType[0] != bashoBanpei {
			fmt.Println()
			fmt.Println()
			bashoBanpei = bashoType[0]

			fmt.Printf("【%s】\n", bashoType[0])
		}

		if bashoType[1] != typeBanpei {
			fmt.Printf(" %s:", bashoType[1])
			typeBanpei = bashoType[1]
		}

		fmt.Print("o")

		// 場所別の集計
		if v, ok := bashoMap[bashoType[0]]; ok {
			bashoMap[bashoType[0]] = v + 1
		} else {
			bashoMap[bashoType[0]] = 1
		}

		// 種類別の集計
		if v, ok := typeMap[bashoType[1]]; ok {
			typeMap[bashoType[1]] = v + 1
		} else {
			typeMap[bashoType[1]] = 1
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(bashoMap)
	fmt.Println()
	fmt.Println(typeMap)
	fmt.Println()
}
