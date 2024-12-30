package main

func main() {

}

type Option interface {
	Get() interface{}
}

func NewSome(val interface{}) Option {

}

type Some struct {
}

func (s *Some) Get() interface{} {
	if s == nil {
		return
	}
}
