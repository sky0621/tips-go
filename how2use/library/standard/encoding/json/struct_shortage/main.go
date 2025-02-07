package main

import (
	"encoding/json"
	"fmt"
)

type Data struct{
	Name string `json:"name"`
}

func (d Data) String() string {
	return fmt.Sprintf("Data{name:%s}", d.Name)
}

var content = `
{
  "name": "Taro",
  "age": 25
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
