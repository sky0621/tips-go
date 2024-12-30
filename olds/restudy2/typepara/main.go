package main

import "fmt"

func main() {
	d1 := &Dir{"d01"}
	f1 := &File{"f01"}

	exec(d1)
	exec(f1)
}

type Request interface {
}
type Response interface {
}

func post[Req Request, Resp Response](req Req, resp Resp) {
	fmt.Println(req)
	fmt.Println(resp)
}

func exec[N Node](n N) {
	n.Describe()
}

var _ Node = (*Dir)(nil)
var _ Node = (*File)(nil)

type Node interface {
	Describe()
}

type Dir struct {
	name string
}

func (d *Dir) Describe() {
	fmt.Println(d.name)
}

type File struct {
	name string
}

func (f File) Describe() {
	fmt.Println(f.name)
}
