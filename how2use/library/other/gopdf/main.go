package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	if err := pdf.AddTTFFont("Ubuntu-L", "./Ubuntu-L.ttf"); err != nil {
		log.Fatal(err)
	}

	if err := pdf.SetFont("Ubuntu-L", "", 14); err != nil {
		log.Fatal(err)
	}

	for _, text := range []string{"text01", "text02", "text03"} {
		pdf.AddPage()
		if err := pdf.Cell(nil, text); err != nil {
			log.Fatal(err)
		}
	}

	if err := pdf.WritePdf("sample.gopdf"); err != nil {
		log.Fatal(err)
	}
}
