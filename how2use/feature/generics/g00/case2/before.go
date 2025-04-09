package main

func MapIntsToStrings(slice []int, f func(int) string) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func MapStringsToInts(slice []string, f func(string) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}
