package main

import (
	"github.com/asticode/go-astilectron"
)

func main() {
	var a, _ = astilectron.New(astilectron.Options{
		AppName:            "MyFsApp",
		AppIconDefaultPath: "def.png",
		AppIconDarwinPath:  "dar.png",
		BaseDirectoryPath:  "/",
	})
	defer a.Close()

	var w, _ = a.NewWindow("http://127.0.0.1:4000", &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(600),
	})
	w.Create()

	a.Start()

	a.Wait()
}
