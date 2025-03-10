package main

import (
	"errors"
	"fmt"
)

var (
	ErrFormat       = errors.New("zip: not a valid zip file")
	ErrAlgorithm    = errors.New("zip: unsupported compression algorithm")
	ErrChecksum     = errors.New("zip: checksum error")
	ErrInsecurePath = errors.New("zip: insecure file path")
)

func main() {
	e := fn(1)
	if errors.Is(e, ErrFormat) {
		fmt.Println(e.Error())
	}
}

func fn(i int) error {
	switch i {
	case 1:
		return ErrFormat
	case 2:
		return ErrAlgorithm
	case 3:
		return ErrChecksum
	case 4:
		return ErrInsecurePath
	case 5:
		return errors.New("other")
	}
	return nil
}
