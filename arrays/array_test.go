package arrays

import "testing"

func BenchmarkValAccess(b *testing.B) {
	for n := 0; n < b.N; n++ {
		intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, val := range intSlice {
			val = val + 2
		}
	}
}

func BenchmarkIndexAccess(b *testing.B) {
	for n := 0; n < b.N; n++ {
		intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for key := range intSlice {
			intSlice[key] = intSlice[key] + 2
		}
	}
}
