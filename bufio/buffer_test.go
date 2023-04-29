package bufio

import "testing"

func BenchmarkOpenCsvWithoutReader(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := OpenCsv("example.csv", false)
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkOpenCsvWithReader(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := OpenCsv("example.csv", true)
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}
