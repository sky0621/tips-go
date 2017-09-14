package main

func main() {

}

type Option interface {
}

type OptionImpl struct {
	data interface{}
	err  error
}
