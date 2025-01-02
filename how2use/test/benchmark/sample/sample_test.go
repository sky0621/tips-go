package sample

import (
	"testing"
	"time"
)

func BenchmarkDoSomething(b *testing.B) {
	time.Sleep(50 * time.Millisecond)
	for i := 0; i < b.N; i++ {
		DoSomething()
	}
}
