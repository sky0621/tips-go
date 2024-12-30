package main

import (
	"fmt"
	"os"

	"github.com/abema/go-mp4"
)

func main() {
	file, err := os.Open("/home/sky0621/ビデオ/tim_williams--slomosf.mp4")
	if err != nil {
		panic(err)
	}
	_, err = mp4.ReadBoxStructure(file, func(h *mp4.ReadHandle) (interface{}, error) {
		fmt.Println("depth", len(h.Path))

		// Box Type (e.g. "mdhd", "tfdt", "mdat")
		fmt.Println(h.BoxInfo.Type.String())

		// Box Size
		fmt.Println(h.BoxInfo.Size)

		// Payload
		if h.BoxInfo.Type.IsSupported() {
			box, _, _ := h.ReadPayload()
			fmt.Println(mp4.Stringify(box))
		}

		// Expands sibling boxes
		return h.Expand()
	})
	if err != nil {
		panic(err)
	}
}
