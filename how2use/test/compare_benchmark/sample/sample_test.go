package sample

import "testing"

// go test -v -count 5 -benchmem -bench BenchmarkBefore 2>&1 | tee old.log
func BenchmarkBefore(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Before(i)
	}
}

// go test -v -count 5 -benchmem -bench BenchmarkAfter 2>&1 | tee new.log
func BenchmarkAfter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = After(i)
	}
}
