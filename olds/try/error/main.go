package main

import "github.com/pkg/errors"

func main() {

}

func a() error {

}

func b() error {

}

func c() error {

}

func d() error {
	err := errors.New()
	return errors.Wrap(err, "at d")
}
