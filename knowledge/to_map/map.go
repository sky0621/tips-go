package main

func toMap[K comparable, T any](slice []T, f func(T) K) map[K]T {
	result := make(map[K]T)
	for _, v := range slice {
		result[f(v)] = v
	}
	return result
}
