package main

func main() {
}

type IoContext interface {
}

type IoContextImpl struct {
}

type InputFunc func() error

type OutputFunc func() error
