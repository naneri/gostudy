package fun

import (
	"math"
	"testing"
)

func BenchmarkPow(b *testing.B) {
	a := 2
	for i := 0; i < b.N; i++ {
		_ = math.Pow(float64(a), 10000)
	}
}

func BenchmarkSelfPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a float64
		a = 2

		for n := 0; n < 1000; n++ {
			a = a * 2
		}
	}
}
