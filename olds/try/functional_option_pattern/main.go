package main

import "context"

func main() {

}

func Get(ctx context.Context, key string, opts ...GetOption) (*GetResponse, error) {

}

type GetResponse struct {
}

type GetOption interface {
	OpOption
}

func Set(ctx context.Context, key, val string, opts ...SetOption) (*SetResponse, error) {

}

type SetResponse struct {
}

type SetOption interface {
	OpOption
}

type OpOption interface {
}
