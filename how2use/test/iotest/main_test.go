package main

import (
	"errors"
	"strings"
	"testing"
	"testing/iotest"
)

// OneByteReaderにより１バイトずつ読み込んだ場合のテスト
func TestOneByteReader(t *testing.T) {
	input := "Hello, OneByteReader!"

	reader := iotest.OneByteReader(strings.NewReader(input))
	output, err := ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	if output != input {
		t.Errorf("expected %q, got %q", input, output)
	}
}

// HalfReaderにより半分ずつ読み込んだ場合のテスト
func TestHalfReader(t *testing.T) {
	input := "Hello, HalfReader!"

	reader := iotest.HalfReader(strings.NewReader(input))
	output, err := ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	if output != input {
		t.Errorf("expected %q, got %q", input, output)
	}
}

func TestErrReader(t *testing.T) {
	reader := iotest.ErrReader(errors.New("error"))
	_, err := ReadAll(reader)
	if err == nil {
		t.Fatal("expected an error, but got nil")
	}
}
