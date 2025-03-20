package main

import (
	"encoding/json"
	"fmt"
)

/*
 * 構造体 B が構造体 A を埋め込んだケースで A が MarshalJSON を実装している場合、
 * A の MarshalJSON が呼ばれ、B の要素は無視される。
 */
func main() {
	b := B{ID: 35, Name: "B", A: A{Name: "A2"}}

	m, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	/*
	 * 期待: {"ID": 35, "Name":"B","Name":"A2"}
	 *
	 * 実際: {"Name":"AAA"}
	 */
	fmt.Println(string(m))

	// 構造体 C が構造体 A を埋め込んだケースで C が MarshalJSON を実装している場合、
	// C の MarshalJSON が呼ばれる。
	c := C{Name: "C", A: A{Name: "A3"}}

	m2, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(m2))
}

type A struct {
	Name string
}

func (a A) MarshalJSON() ([]byte, error) {
	return []byte(`{"Name": "AAA"}`), nil
}

type B struct {
	ID   int
	Name string
	A
}

type C struct {
	Name string
	A
}

func (c C) MarshalJSON() ([]byte, error) {
	m, err := c.A.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf(`{"Name": "%s", "A":%s}`, c.Name, string(m))), nil
}
