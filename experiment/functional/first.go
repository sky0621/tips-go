package functional

// First ...
func First(ary []int) int {
	if ary == nil || len(ary) == 0 {
		return 0
	}
	return ary[0]
}
