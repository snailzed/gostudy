package util

import "testing"

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(16)
	}
}
