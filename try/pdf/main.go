package main

import (
	"flag"
	"fmt"
	"github.com/signintech/gopdf"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	if err := pdf.AddTTFFont("Ubuntu-L", "./Ubuntu-L.ttf"); err != nil {
		log.Fatal(err)
	}

	if err := pdf.SetFont("Ubuntu-L", "", 14); err != nil {
		log.Fatal(err)
	}

	pdf.AddPage()
	pdf.Cell(nil, "A1")
	pdf.AddPage()
	pdf.Cell(nil, "A2")

	pdf.WritePdf("sample.pdf")
}
