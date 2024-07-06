package inlinex

import (
	"testing"
)

//go:noinline
func maxNoinline(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxInline(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// BenchmarkNoInline-10    	576585926	         2.067 ns/op
func BenchmarkNoInline(b *testing.B) {
	x, y := 1, 2
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxNoinline(x, y)
	}
}

// BenchmarkInline-10    	1000000000	         0.3176 ns/op
func BenchmarkInline(b *testing.B) {
	x, y := 1, 2
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxInline(x, y)
	}
}
