package sample

import "strconv"

func Simple(i int) string {
	return strconv.Itoa(i)
}

func Complex(i int) string {
	switch i {
	case 0:
		return "zero"
	case 1:
		return "plus"
	case -1:
		return "minus"
	}
	return ""
}
