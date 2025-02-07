package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
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
  },
  "fly": true,
  "sights": [
    "2023-01-17",
    "2023-08-30",
    "2024-05-24"
  ],
  "pt": 4.05
}
`

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (d Dimensions) String() string {
	return fmt.Sprintf("{width: %d, height: %d}", d.Width, d.Height)
}

type Sights []string

func (s Sights) String() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	for i, v := range s {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(v)
	}
	sb.WriteString("]")
	return sb.String()
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
	Fly         bool       `json:"fly"`
	Sights      Sights     `json:"sights"`
	Pt          float32    `json:"pt"`
}

func (d Data) String() string {
	return fmt.Sprintf("{species: %s, description: %s, dimensions: %s, fly: %v, sights: %s, pt: %v}", d.Species, d.Description, d.Dimensions.String(), d.Fly, d.Sights.String(), d.Pt)
}
