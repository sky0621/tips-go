package main

import "fmt"

// image
// func fn(b bool)(i int)(s string) string {}

type StrToStr func(s string) string

type IntToStrToStr func(i int) StrToStr

type Fn func(b bool) IntToStrToStr

func fn(b bool) IntToStrToStr {
	return func(i int) StrToStr {
		return func(s string) string {
			if b {
				if i > 0 {
					return fmt.Sprintf("[T][+][%s]", s)
				} else {
					return fmt.Sprintf("[T][-][%s]", s)
				}
			} else {
				if i > 0 {
					return fmt.Sprintf("[F][+][%s]", s)
				} else {
					return fmt.Sprintf("[F][-][%s]", s)
				}
			}
		}
	}
}
