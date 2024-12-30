package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data Data
	if err := json.Unmarshal([]byte(content), &data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

var content = `
{
  "species": "ハト",
  "description": "岩にとまるのが好き",
  "dimensions": {
    "height": 24,
    "width": 10
  }
}
`

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (d Dimensions) String() string {
	return fmt.Sprintf("{width: %d, height: %d}", d.Width, d.Height)
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
}

func (d Data) String() string {
	return fmt.Sprintf("{species: %s, description: %s, dimensions: %s}", d.Species, d.Description, d.Dimensions.String())
}
