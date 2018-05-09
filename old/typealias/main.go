package main

import "fmt"

func main() {
	var f foo
	var b bar

	b = bar{}
	f.say("foo")
	b.say("bar")
	//fmt.Println(b)

	var b2 bar
	b2 = bar(f)
	b2.say("bar2")
	b2.say2("bar2bar2")
}

type foo struct {
}

func (f *foo) say(w string) {
	fmt.Println(w)
}

//type bar foo
type bar = foo

func (b *bar) say2(w string) {
	fmt.Printf("say2:%v",w)
}