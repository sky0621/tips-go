package main

import "log"

func main() {
	apple := Fruit{name: "Apple"}
	res01 := apple.show("Yummy!")
	orange := Fruit{name: "Orange"}
	res02 := orange.show("Soso.")
	log.Println(res01, apple.name)
	log.Println(res02, orange.name)
	log.Println("---------------------------------------------")
	aMelon := &Fruit{name: "Melon"}
	res03 := aMelon.show("green")
	aPeach := &Fruit{name: "Peach"}
	res04 := aPeach.show("pink...")
	log.Println(res03, aMelon.name)
	log.Println(res04, aPeach.name)
	aPeach.name = "NotPeach"
	res04b := aPeach.show("pink...")
	log.Println(res04b, aPeach.name)
	log.Println("---------------------------------------------")
	log.Println("---------------------------------------------")
	log.Println("---------------------------------------------")
	apple2 := Fruit{name: "Apple"}
	res012 := apple2.show2("Yummy!")
	orange2 := Fruit{name: "Orange"}
	res022 := orange2.show2("Soso.")
	log.Println(res012, apple2.name)
	log.Println(res022, orange2.name)
	log.Println("---------------------------------------------")
	aMelon2 := &Fruit{name: "Melon"}
	res032 := aMelon2.show2("green")
	aPeach2 := &Fruit{name: "Peach"}
	res042 := aPeach2.show2("pink...")
	log.Println(res032, aMelon2.name)
	log.Println(res042, aPeach2.name)
	aPeach2.name = "NotPeach"
	res04b2 := aPeach2.show2("pink...")
	log.Println(res04b2, aPeach2.name)
	log.Println("---------------------------------------------")
}

// Fruit ...
type Fruit struct {
	name string
}

func (f Fruit) show(str string) string {
	f.name = "[" + f.name + "]"
	return f.name + " is " + str
}

func (f *Fruit) show2(str string) string {
	f.name = "[" + f.name + "]"
	return f.name + " is " + str
}
