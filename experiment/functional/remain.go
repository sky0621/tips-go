package functional

// Remain ...
func Remain(ary []int) []int {
	if ary == nil || len(ary) <= 1 {
		return []int{0}
	}
	return ary[1:]
}
