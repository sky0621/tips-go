package main

import (
	"encoding/json"
	"fmt"
)

type Data struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func (d Data) String() string {
	return fmt.Sprintf("Data{name:%s, age:%d}", d.Name, d.Age)
}

var content = `
{
  "name": "Taro"
}
`

func main() {
	var data Data
	if err := json.Unmarshal([]byte(content), &data);err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
