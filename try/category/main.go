package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	o1 := NewObject()
	o2 := NewObject()

	o1.Add(NewArrow(o1, o2))

	c := NewCategory([]Object{o1, o2})
	fmt.Println(c.IsMonoid())
}

func NewCategory(objects []Object) Category {
	return &category{
		objects: objects,
	}
}

// 圏
type Category interface {
	IsMonoid() bool
}

type category struct {
	objects []Object
}

// 対象がただ1つの圏をモノイドと言う
func (c *category) IsMonoid() bool {
	return c.objects != nil && len(c.objects) == 1
}

func NewObject() Object {
	o := &object{}
	// 恒等射は必ず存在
	identity := NewArrow(o, o)
	o.arrows = append(o.arrows, identity)
	return o
}

// 対象
type Object interface {
	Identity() Arrow
	Add(a Arrow) bool
	Adds(as []Arrow) bool
}

type object struct {
	arrows []Arrow
}

// 恒等射
func (o *object) Identity() Arrow {
	return o.arrows[0]
}

func (o *object) Add(a Arrow) bool {
	if a == nil {
		return false
	}
	o.arrows = append(o.arrows, a)
	return true
}

func (o *object) Adds(as []Arrow) bool {
	if as == nil {
		return false
	}
	o.arrows = append(o.arrows, as...)
	return true
}

func NewArrow(domain, codomain Object) Arrow {
	return &arrow{
		domain:   domain,
		codomain: codomain,
	}
}

// 射
type Arrow interface {
}

type arrow struct {
	domain   Object
	codomain Object
}
