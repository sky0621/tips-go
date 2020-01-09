package main

func main() {

}

func NewCategory() Category {
	return &category{}
}

type Category interface {
}

type category struct {
	objects []Object
}

func NewObject() Object {
	return &object{}
}

type Object interface {
}

type object struct {
}

func NewArrow(domain, codomain Object) Arrow {
	return &arrow{
		domain:   domain,
		codomain: codomain,
	}
}

type Arrow interface {
	Domain() Object
	Codomain() Object
}

type arrow struct {
	domain   Object
	codomain Object
}

func (a *arrow) Domain() Object {
	if a == nil {
		return nil
	}
	return a.domain
}

func (a *arrow) Codomain() Object {
	if a == nil {
		return nil
	}
	return a.codomain
}
