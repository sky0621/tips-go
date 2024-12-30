package fp

// Sum ...
// [1, 2, 3]
// 1 + Sum([2, 3])
// 1 + 2 + Sum([3])
func Sum(ary []int) int {
	if ary == nil || len(ary) == 0 {
		return 0
	}
	if len(ary) == 1 {
		return First(ary)
	}
	return First(ary) + Sum(Remain(ary))
}
