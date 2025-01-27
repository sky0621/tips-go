package person

func isID(v int64) bool {
	return v > 0
}

func isName(v string) bool {
	return v != ""
}
