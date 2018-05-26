package fp

// Add ... [x + 1], [y - 1] repeat (if y > 0)
func Add(x, y int) int {
	if y < 1 {
		return x
	}
	return Add(x+1, y-1)
}
